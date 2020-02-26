package main

import "fmt"

type Intf interface {
	M()
	N()
}

func main() {
	t1 := T{"无用"}
	fmt.Println(t1.Name)
	t1.M()
	fmt.Println(t1.Name)
	t1.N()
	fmt.Println(t1.Name)
}

type T struct {
	Name string
}

func (t T) M() {
	t.Name = "小明"
}

func (t *T) N() {
	t.Name = "红"
}
