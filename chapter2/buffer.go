package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	//bufferにためておく
	buffer.Write([]byte("bytes.Buffer example1\n"))
	buffer.Write([]byte("bytes.Buffer example2\n"))
	fmt.Println(buffer.String())
}
