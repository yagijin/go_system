package main

import (
	"os"
)

func main() {
	//標準出力
	os.Stdout.Write([]byte("os.Stdout example\n"))

	//標準エラー出力
	os.Stderr.Write([]byte("os.Stderr example\n"))
}
