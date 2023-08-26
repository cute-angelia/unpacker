package unpacker

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
)

var zipHeader = []byte("PK\x03\x04")

// https://py7zr.readthedocs.io/en/latest/archive_format.html#signature
var sevenZipHeader = []byte("7z\xBC\xAF\x27\x1C")
var (
	rarHeaderV1_5 = []byte("Rar!\x1a\x07\x00")     // v1.5
	rarHeaderV5_0 = []byte("Rar!\x1a\x07\x01\x00") // v5.0
)

func GetUnpacker(filepath string) (Packer, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// check zip
	if v, err := isZip(filepath, f); err == nil {
		return v, nil
	}

	// check 7z
	if v, err := is7Z(filepath, f); err == nil {
		return v, nil
	}

	isRar(filepath, f)
	isRar2(filepath, f)

	return nil, errors.New("文件类型不支持")
}

func isZip(filepath string, f io.Reader) (Packer, error) {
	buf, err := readAtMost(f, len(zipHeader))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(buf, zipHeader) {
		log.Println("is zip")
		return NewZip(filepath), nil
	} else {
		return nil, errors.New("not zip")
	}
}

func is7Z(filepath string, f io.Reader) (Packer, error) {
	buf, err := readAtMost(f, len(sevenZipHeader))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(buf, sevenZipHeader) {
		log.Println("is 7z")
		return NewSevenZip(filepath), nil
	} else {
		return nil, errors.New("not 7z")
	}
}

func isRar(filepath string, f io.Reader) (Packer, error) {
	buf, err := readAtMost(f, len(rarHeaderV5_0))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(buf, rarHeaderV5_0) {
		log.Println("is rar5")
		return NewSevenZip(filepath), nil
	} else {
		return nil, errors.New("not rar")
	}
}
func isRar2(filepath string, f io.Reader) (Packer, error) {
	buf, err := readAtMost(f, len(rarHeaderV1_5))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(buf, rarHeaderV1_5) {
		log.Println("is rar1")
		return NewSevenZip(filepath), nil
	} else {
		return nil, errors.New("not rar")
	}
}
