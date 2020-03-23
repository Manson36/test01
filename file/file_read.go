package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//O_...中是二进制的代码，每一位表示不同的功能，使用逻辑或('|')表任何一个功能实现
	if err != nil {
		fmt.Println("open file failed, err", err)
		return
	}
	defer fileObj.Close()

	//方法一：
	// write
	fileObj.Write([]byte("linqintaimeile\n"))
	//write string
	fileObj.WriteString("白鹿也很美")
	//方法二：bufio
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello, 女神") //写到缓存中
	wr.Flush()                  //将缓存文件内容写入文件
	//方法三：ioutil
	writeFileByioutil()
}

func writeFileByioutil() {
	str := "hello, 女神！"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err", err)
		return
	}
}
