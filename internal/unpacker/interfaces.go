package unpacker

type Packer interface {
	Name() string
	Extract(targetDir string) error
}
