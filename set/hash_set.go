package set

import "sync"

type HashSet struct {
	mu  sync.Mutex
	set map[interface{}]Item
}

func NewHashSet() *HashSet {
	return &HashSet{
		set: make(map[interface{}]Item),
	}
}

func (hs *HashSet) Add(item Item) {
	hs.mu.Lock()
	hs.set[item.Hash()] = item
	hs.mu.Unlock()
}

func (hs *HashSet) AddAll(items []Item) {
	hs.mu.Lock()
	for _, item := range items {
		hs.set[item.Hash()] = item
	}
	hs.mu.Unlock()
}

func (hs *HashSet) Clear() {
	hs.mu.Lock()
	hs.set = make(map[interface{}]Item)
	hs.mu.Unlock()
}

func (hs *HashSet) Contains(item Item) bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	if _, ok := hs.set[item.Hash()]; ok {
		return true
	}

	return false
}

func (hs *HashSet) IsEmpty() bool {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	if len(hs.set) > 0 {
		return false
	}

	return true
}

func (hs *HashSet) Remove(item Item) {
	hs.mu.Lock()
	delete(hs.set, item.Hash())
	hs.mu.Unlock()
}

func (hs *HashSet) RemoveAll(items []Item) {
	hs.mu.Lock()
	for _, item := range items {
		delete(hs.set, item.Hash())
	}
	hs.mu.Unlock()
}

func (hs *HashSet) Size() int {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	return len(hs.set)
}

func (hs *HashSet) Slice() []Item {
	hs.mu.Lock()
	slice := make([]Item, 0, len(hs.set))
	for _, item := range hs.set {
		slice = append(slice, item)
	}
	hs.mu.Unlock()

	return slice
}
