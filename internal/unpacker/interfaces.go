package unpacker

type Packer interface {
	Name() string
	Extract() error
}
