package main

import (
	"fmt"
)

func main() {
	//var s string
	//fmt.Scan(&s)
	//fmt.Println("s的值是：", s)

	var name string
	var age int
	var class string
	//fmt.Scanf("%s %d %s\n", &name, &age, &class) //这里引号中就不要多此一举加逗号了，而且要有'\n'
	//fmt.Println(name, age, class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
