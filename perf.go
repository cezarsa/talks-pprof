package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sort"
)

type orderedList struct {
	ordered []int
}

func (l *orderedList) insert(val int) {
	pos := sort.SearchInts(l.ordered, val)
	l.ordered = append(l.ordered, val)
	copy(l.ordered[pos+1:], l.ordered[pos:])
	l.ordered[pos] = val
}

func (l *orderedList) items() []int {
	return l.ordered
}

func main() {
	list := orderedList{}
	nCh := make(chan int)
	go func() {
		for n := range nCh {
			list.insert(n)
		}
	}()
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var n int
			fmt.Fscanf(r.Body, "%d", &n)
			r.Body.Close()
			nCh <- n
			return
		}
		fmt.Fprintf(w, "%#v", list.items())
	}))
	http.ListenAndServe(":8000", nil)
}
