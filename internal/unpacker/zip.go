package unpacker

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
)

type Zip struct {
	filepath string
}

func NewZip(filepath string) *Zip {
	return &Zip{
		filepath: filepath,
	}
}

func (z Zip) Name() string { return ".zip" }

func (z Zip) Extract(targetDir string) error {
	archive := z.filepath

	r, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	Encryption := false

	defer r.Close()
	for _, f := range r.File {

		if f.Flags&0x1 == 1 {
			Encryption = true
		} else {
			Encryption = false
		}
		if Encryption {
			log.Println("已加密")
			return errors.New("已加密")
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		destFile, _ := os.Create(targetDir + "/" + f.Name)
		// srcFile复制到destFile
		io.Copy(destFile, rc)
		rc.Close()
		destFile.Close()
	}
	return nil
}
