package set

import (
	"container/list"
	"sync"
)

type OrderedHashSet struct {
	mu sync.Mutex
	HashSet
	list *list.List
}

func NewOrderedHashSet() *OrderedHashSet {
	return &OrderedHashSet{
		HashSet: *NewHashSet(),
		list:    list.New(),
	}
}

func (hs *OrderedHashSet) Add(item Item) {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	if hs.HashSet.Contains(item) {
		return
	}

	hs.HashSet.Add(item)
	hs.list.PushBack(item)
}

func (hs *OrderedHashSet) AddAll(items ...Item) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	for _, item := range items {
		if hs.HashSet.Contains(item) {
			continue
		}

		hs.HashSet.Add(item)
		hs.list.PushBack(item)
	}
}

func (hs *OrderedHashSet) Clear() {
	hs.mu.Lock()
	hs.HashSet.Clear()
	hs.list.Init()
	hs.mu.Unlock()
}

func (hs *OrderedHashSet) Remove(item Item) {
	hs.mu.Lock()
	hs.Remove(item)
	var curr *list.Element
	for curr = hs.list.Front(); curr != nil; curr = curr.Next() {
		found, ok := curr.Value.(Item)
		if !ok {
			panic("set: unable to assert Item")
		}

		if found == item {
			break
		}
	}

	if curr != nil {
		hs.list.Remove(curr)
	}

	hs.mu.Unlock()
}

func (hs *OrderedHashSet) RemoveAll(items ...Item) {
	hs.mu.Lock()
	for _, item := range items {
		hs.Remove(item)
		var curr *list.Element
		for curr = hs.list.Front(); curr != nil; curr = curr.Next() {
			found, ok := curr.Value.(Item)
			if !ok {
				panic("set: unable to assert Item")
			}

			if found == item {
				break
			}
		}

		if curr != nil {
			hs.list.Remove(curr)
		}
	}
	hs.mu.Unlock()
}

func (hs *OrderedHashSet) Slice() []Item {
	hs.mu.Lock()
	slice := make([]Item, 0, hs.HashSet.Size())
	for curr := hs.list.Front(); curr != nil; curr = curr.Next() {
		item, ok := curr.Value.(Item)
		if !ok {
			panic("set: unable to assert Item")
		}

		slice = append(slice, item)
	}
	hs.mu.Unlock()

	return slice
}
