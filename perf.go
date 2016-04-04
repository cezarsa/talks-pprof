package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sort"
	"strconv"
	"sync"
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

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 10)
	},
}

func numBin(w io.Writer) {
	data := pool.Get().([]byte)
	for i := int64(0); i < 1000; i++ {
		data = strconv.AppendInt(data, rand.Int63n(1000), 2)
		data = append(data, '\n')
	}
	w.Write(data)
	pool.Put(data[:0])
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
	http.Handle("/buffer", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numBin(w)
	}))
	http.ListenAndServe(":8000", nil)
}
