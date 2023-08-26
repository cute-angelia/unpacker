package unpacker

import (
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unpacker/internal/password"
)

type Zip struct {
	filePath    string
	target      string
	password    string
	count       int // 重试次数
	isEncrypted bool
}

func NewZip(filePath string) *Zip {
	// GetFileName 根据路径获取文件名
	ext := path.Ext(filePath)
	target := path.Dir(filePath) + "/" + path.Base(filePath[:len(filePath)-len(ext)])

	return &Zip{
		filePath: filePath,
		target:   target,
		count:    0,
	}
}

func (z Zip) Name() string { return ".zip" }

func (z Zip) Extract() error {
	r, err := zip.OpenReader(z.filePath)
	if err != nil {
		return err
	}

	defer r.Close()

	passwords := password.GetPasswords()
	maxTimes := len(passwords)

	os.MkdirAll(z.target, 0777)

label:
	for _, f := range r.File {
		if f.IsEncrypted() {
			z.isEncrypted = true
			log.Println("已加密, 读取密码", z.password, "次数", z.count)
			f.SetPassword(z.password)
		}

		fullPath := filepath.Join(z.target, f.Name)

		if strings.Contains(fullPath, "__MACOSX") {
			continue
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(fullPath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}
		log.Println("解压", f.Name)
		fileReader, err := f.Open()
		if err != nil {
			log.Println("reader", err)
			return err
		}

		targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Println(err)
			fileReader.Close()
			return err
		}
		_, err = io.Copy(targetFile, fileReader)
		// close all
		fileReader.Close()
		targetFile.Close()
		if err != nil {
			log.Println(err)
			z.password = passwords[z.count]
			z.count++
			if z.count >= maxTimes {
				break
			}
			goto label
		}
	}
	return nil
}
