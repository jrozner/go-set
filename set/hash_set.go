package set

import "sync"

// HashSet is a set implementation that uses the Go map as the underlying
// storage mechanism. The Item's Hash method is called and the resulting
// value used to denote the Item's value in the set. HashSet is unordered
type HashSet struct {
	mu  sync.Mutex
	set map[interface{}]Item
}

// NewHashSet returns a pointer to a new HashSet
func NewHashSet() *HashSet {
	return &HashSet{
		set: make(map[interface{}]Item),
	}
}

// Add adds the item to the set
func (hs *HashSet) Add(item Item) {
	hs.mu.Lock()
	hs.set[item.Hash()] = item
	hs.mu.Unlock()
}

// AddAll adds all items to the set
func (hs *HashSet) AddAll(items ...Item) {
	hs.mu.Lock()
	for _, item := range items {
		hs.set[item.Hash()] = item
	}
	hs.mu.Unlock()
}

// Clear removes all items from the set
func (hs *HashSet) Clear() {
	hs.mu.Lock()
	hs.set = make(map[interface{}]Item)
	hs.mu.Unlock()
}

// Contains returns a boolean denoting whether the set contains the specified
// Item or not
func (hs *HashSet) Contains(item Item) bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	if _, ok := hs.set[item.Hash()]; ok {
		return true
	}

	return false
}

// IsEmpty returns a boolean denoting whether the set contains any Items
func (hs *HashSet) IsEmpty() bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	if len(hs.set) > 0 {
		return false
	}

	return true
}

// Remove removes the specified Item from the set
func (hs *HashSet) Remove(item Item) {
	hs.mu.Lock()
	delete(hs.set, item.Hash())
	hs.mu.Unlock()
}

// RemoveAll removes the specified Items from the set
func (hs *HashSet) RemoveAll(items ...Item) {
	hs.mu.Lock()
	for _, item := range items {
		delete(hs.set, item.Hash())
	}
	hs.mu.Unlock()
}

// Size returns an integer specifying the number of Items in the set
func (hs *HashSet) Size() int {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	return len(hs.set)
}

// Slice returns a []Item of all the items in the set. In this implementation
// the slice's order will not be preserved across calls to this method and
// the order will not represent the order of the items added to the set
func (hs *HashSet) Slice() []Item {
	hs.mu.Lock()
	slice := make([]Item, 0, len(hs.set))
	for _, item := range hs.set {
		slice = append(slice, item)
	}
	hs.mu.Unlock()

	return slice
}
