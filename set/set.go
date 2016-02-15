package set

// Set is an interface for all Set implementations to adhere to
type Set interface {
	Add(Item)
	AddAll(...Item)
	Clear()
	Contains(Item) bool
	IsEmpty() bool
	Remove(Item)
	RemoveAll(...Item)
	Size() int
	Slice() []Item
}
