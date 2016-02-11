package set

type Set interface {
	Add(interface{})
	AddAll([]interface{})
	Clear()
	Contains(interface{}) bool
	IsEmpty() bool
	Remove(interface{})
	RemoveAll([]interface{})
	Size() int
	Slice() []interface{}
}
