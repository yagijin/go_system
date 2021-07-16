package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings.Builderは書き出し専用のbytes.Buffer
	//なので読み出しはbyteではなくStringしかできない
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
