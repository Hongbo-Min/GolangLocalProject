package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { // strings.HasPrefix 用来检查字符串是否以给定字符开始
			url = "http://" + url
		}
		fmt.Println("url: ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("httpStatus: ", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body) // io.Copy可以避免申请一个缓冲区来存储
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
