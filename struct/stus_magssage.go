package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student

func showAllStudent() {
	for i, v := range allStudent {
		fmt.Printf("学号：%d，姓名：%v\n", i, v.name)
	}
}

func addStudent() {
	var id int64
	var name string
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
}

func main() {
	allStudent = make(map[int64]*student, 50)
	for {
		fmt.Println("欢迎光临数字管理系统")
		fmt.Println("1.查看所有学生 2.新增学生 3.删除学生 4.退出")

		fmt.Print("请输入你要干啥？")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你学则的选项是：", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStrudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请重新输入")
		}
	}
}
