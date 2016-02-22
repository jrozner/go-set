package main

import "github.com/jrozner/go-set/set"

type CoolThing struct {
	s *set.SimpleSet
}

func NewCoolThing() *CoolThing {
	return &CoolThing{
		s: set.NewSimpleSet(),
	}
}

func (s *CoolThing) Add(item set.Item) {
	s.s.Add(item)
}
