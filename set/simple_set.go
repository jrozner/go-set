package set

import "sync"

// SimpleSet is an unordered Set that is implemented by using the Item itself
// as the key in a Go hash. If more complex logic is desired for determining
// the value HashSet should be used intead which allows the use of the Hash
// method to determine the uniqueness of the Item in the set. This is the
// efficient, in terms of allocations and memory use, of the set
// implementations
type SimpleSet struct {
	mu  sync.Mutex
	set map[Item]struct{}
}

// NewSimpleSet returns a pointer to a new SimpleSet
func NewSimpleSet() *SimpleSet {
	return &SimpleSet{
		set: make(map[Item]struct{}),
	}
}

// Add adds the item to the set
func (hs *SimpleSet) Add(item Item) {
	hs.mu.Lock()
	hs.set[item] = struct{}{}
	hs.mu.Unlock()
}

// AddAll adds all items to the set
func (hs *SimpleSet) AddAll(items ...Item) {
	hs.mu.Lock()
	for _, item := range items {
		hs.set[item] = struct{}{}
	}
	hs.mu.Unlock()
}

// Clear removes all items from the set
func (hs *SimpleSet) Clear() {
	hs.mu.Lock()
	hs.set = make(map[Item]struct{})
	hs.mu.Unlock()
}

// Contains returns a boolean denoting whether the set contains the specified
// Item or not
func (hs *SimpleSet) Contains(item Item) bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	if _, ok := hs.set[item]; ok {
		return true
	}

	return false
}

// IsEmpty returns a boolean denoting whether the set contains any Items
func (hs *SimpleSet) IsEmpty() bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	if len(hs.set) > 0 {
		return false
	}

	return true
}

// Remove removes the specified Item from the set
func (hs *SimpleSet) Remove(item Item) {
	hs.mu.Lock()
	delete(hs.set, item)
	hs.mu.Unlock()
}

// RemoveAll removes the specified Items from the set
func (hs *SimpleSet) RemoveAll(items ...Item) {
	hs.mu.Lock()
	for _, item := range items {
		delete(hs.set, item)
	}
	hs.mu.Unlock()
}

// Size returns an integer specifying the number of Items in the set
func (hs *SimpleSet) Size() int {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	return len(hs.set)
}

// Slice returns a []Item of all the items in the set. In this implementation
// the slice's order will not be preserved across calls to this method and
// the order will not represent the order of the items added to the set
func (hs *SimpleSet) Slice() []Item {
	hs.mu.Lock()
	slice := make([]Item, 0, len(hs.set))
	for item := range hs.set {
		slice = append(slice, item)
	}
	hs.mu.Unlock()

	return slice
}
