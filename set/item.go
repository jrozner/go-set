package set

type Item interface {
	Hash() interface{}
	Value() interface{}
}
