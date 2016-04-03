package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestOrderedListInsert(t *testing.T) {
	a := []int{9, 1, 7, 5, 10, 20, 15, 5}
	list := orderedList{}
	for _, v := range a {
		list.insert(v)
	}
	if !reflect.DeepEqual(list.items(), []int{1, 5, 5, 7, 9, 10, 15, 20}) {
		t.Fatalf("got %#v", list.items())
	}
}

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
