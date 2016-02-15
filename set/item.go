package set

// Item is a container for storing values in Sets
type Item interface {
	Hash() interface{}
	Value() interface{}
}
