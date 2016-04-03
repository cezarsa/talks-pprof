package main

import (
	"bufio"
	"os"
	"runtime/pprof"
	"strconv"
)

type orderedList struct {
	root *listEntry
}

type listEntry struct {
	val  int
	next *listEntry
}

func (l *orderedList) insert(val int) {
	el := l.root
	var prevEl *listEntry
	for el != nil && val > el.val {
		prevEl = el
		el = el.next
	}
	newEl := &listEntry{val: val, next: el}
	if prevEl == nil {
		l.root = newEl
	} else {
		prevEl.next = newEl
	}
}

func (l *orderedList) items() []int {
	var items []int
	el := l.root
	for el != nil {
		items = append(items, el.val)
		el = el.next
	}
	return items
}

func addFromStdin() {
	list := orderedList{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.ParseInt(line, 10, 32)
		list.insert(int(n))
	}
}

func main() {
	file, err := os.Create("cpu_profile.out")
	if err != nil {
		panic("unable to open file: " + err.Error())
	}
	pprof.StartCPUProfile(file)
	defer file.Close()
	defer pprof.StopCPUProfile()
	if os.Args[1] == "add" {
		addFromStdin()
	}
}
