package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//ファイル出力
	file.Write([]byte("os.File example\n"))
	file.Close()
}
