package main

import (
	"math/rand"
	"testing"
)

func BenchmarkOrderedListInsert(b *testing.B) {
	list := orderedList{}
	for i := 0; i < b.N; i++ {
		list.insert(rand.Int())
	}
	b.StopTimer()
	items := list.items()
	var lastItem int
	for i, item := range items {
		if i > 0 && item < lastItem {
			b.Fatalf("items not ordered: %#v", items)
		}
		lastItem = item
	}
}
