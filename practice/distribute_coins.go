package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：
Matthew，Sarah，Augustus，Heidi，Emilie，Peter，Giana，Adriano，Aaron，Elizabeth。
分配规则如下：
a.名字中每包含一个'e'或'E'分1枚金币
b.名字中每包含一个'i'或'I'分2枚金币
c.名字中每包含一个'o'或'O'分3枚金币
d.名字中每包含一个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
*/

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	//1. 依次拿到每个人的名字
	//2. 拿到一个人名，根据规则分配金币，保存到distribution中，还要记录剩下的金币
	for _, name := range users {
		sum := 0
		for _, v := range name {
			switch v {
			case 'e', 'E':
				sum += 1
			case 'i', 'I':
				sum += 2
			case 'o', 'O':
				sum += 3
			case 'u', 'U':
				sum += 4
			}
		}
		distribution[name] = sum
		fmt.Printf("%s 获得金币 %d\n", name, sum)
		coins -= sum
	}
	return coins
}
