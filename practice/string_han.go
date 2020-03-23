package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中有几个汉字
	var s1 = "hello李沁"
	var count = 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) { //判断是否是汉字
			count++
		}
	}
	fmt.Println(count)

	//2.判断一段英文中有单词出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ") //把字符串按空格切割成字符串切片

	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; ok { //判断单词在map中是否存在
			m1[w] += 1 //单词在map中存在数量加一
		} else {
			m1[w] = 1 //不存在，值为一
		}
	}
	for i, v := range m1 {
		fmt.Println(i, v)
	}

	//3.回文判断：例子 山西运煤车煤运西山；黄山落叶松叶落山黄
	ss := "黄山落叶松叶落山黄"
	r := make([]rune, 0, len(ss)) //将字符串中的字符放入[]rune中
	for _, s := range ss {
		r = append(r, s)
	}

	for i := 0; i < len(r)/2; i++ {
		if r[0] != r[len(r)-1] {
			fmt.Println("不是回文")
			break
		}
	}
	fmt.Println("是回文")
}
