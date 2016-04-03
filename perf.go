package main

import (
	"bufio"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
)

type orderedList struct {
	ordered []int
}

func (l *orderedList) insert(val int) {
	pos := sort.SearchInts(l.ordered, val)
	list := append([]int{}, l.ordered[:pos]...)
	list = append(list, val)
	l.ordered = append(list, l.ordered[pos:]...)
}

func (l *orderedList) items() []int {
	return l.ordered
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
