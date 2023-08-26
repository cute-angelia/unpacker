package unpacker

import (
	"github.com/bodgit/sevenzip"
	"log"
	"path"
)

type SevenZip struct {
	// If true, errors encountered during reading or writing
	// a file within an archive will be logged and the
	// operation will continue on remaining files.
	ContinueOnError bool

	// The password, if dealing with an encrypted archive.
	Password string

	filePath    string
	target      string
	password    string
	count       int // 重试次数
	isEncrypted bool
}

func NewSevenZip(filePath string) *SevenZip {
	// GetFileName 根据路径获取文件名
	ext := path.Ext(filePath)
	target := path.Dir(filePath) + "/" + path.Base(filePath[:len(filePath)-len(ext)])

	return &SevenZip{
		filePath: filePath,
		target:   target,
	}
}

func (z SevenZip) Name() string { return ".7z" }

// Extract extracts files from z, implementing the Extractor interface. Uniquely, however,
// sourceArchive must be an io.ReaderAt and io.Seeker, which are oddly disjoint interfaces
// from io.Reader which is what the method signature requires. We chose this signature for
// the interface because we figure you can Read() from anything you can ReadAt() or Seek()
// with. Due to the nature of the zip archive format, if sourceArchive is not an io.Seeker
// and io.ReaderAt, an error is returned.
func (z SevenZip) Extract() error {
	r, err := sevenzip.OpenReaderWithPassword(z.filePath, z.password)
	if err != nil {
		log.Println(err)
	}
	defer r.Close()
	return nil
}
