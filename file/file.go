package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件
	fileObj, err := os.Open("./file/file.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()

	//readFromFile1(fileObj)
	readFromFielByBufio(fileObj)
	//readFromFileByioutil()
}

func readFromFile1(fileObj *os.File) {
	//读取文件
	var tmp = make([]byte, 128) //指定读取的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Printf("读取%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < 128 {
			return
		}
	}
}

func readFromFielByBufio(fileObj *os.File) {
	//读取文件，方法二：使用bufio读取
	//创建从文件中读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //按行读取
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file failed, err", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByioutil() {
	//读取文件，方法三：使用ioutil读取全部文件
	ret, err := ioutil.ReadFile("./file/file.go")
	if err != nil {
		fmt.Println("read fiel failed, err:", err)
		return
	}
	fmt.Println(string(ret))
}
