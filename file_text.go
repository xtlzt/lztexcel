package main

import (
	"fmt"
	"lztexcel/work"
)

func main() {
	file := work.FileInfo("C:/Users/NINGMEI/Desktop/需要的信息.txt", 1048)
	fmt.Println(file.FileExt)
	fmt.Println(file.FileMode)
	fmt.Println(file.FileModTime)
	fmt.Println(file.Sys)
}
