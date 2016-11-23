package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	if p[i] > p[j] {
		return false
	} else {
		return true
	}
}

func (p people) Swap(i, j int) {
	var v string
	v = p[i]
	p[i] = p[j]
	p[j] = v
}

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	studyGroup = people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

	fmt.Println(n)
	sort.Ints(n)
	sort.Reverse(sort.IntSlice(n))
	fmt.Println(n)

}

// Use https://godoc.org/sort to sort the following:
//
// (1)
// type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}
//
// (2)
// s := []string{"Zeno", "John", "Al", "Jenny"}
//
// (3)
// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
//
// Also sort the above in reverse order
