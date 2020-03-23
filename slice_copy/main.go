package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 2, 3}
	a3 := make([]int, 3, 3)
	copy(a3, a1)
	a1[0] = 1000
	fmt.Println(a1, a3)
	sort.Ints(a1)
	fmt.Println(a1, a3)
}
