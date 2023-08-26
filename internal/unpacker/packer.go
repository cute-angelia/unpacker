package unpacker

import (
	"bytes"
	"os"
)

var zipHeader = []byte("PK\x03\x04")

func GetUnpacker(filepath string) (Packer, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := readAtMost(f, len(zipHeader))
	if err != nil {
		return nil, err
	}
	if bytes.Equal(buf, zipHeader) {
		return NewZip(filepath), nil
	}
	return nil, err
}
