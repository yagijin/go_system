package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-Test", "ヘッダーお追加できます")
	request.Write(os.Stdout)
}
