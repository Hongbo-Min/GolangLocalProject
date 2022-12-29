package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 从标准输入读取重复的行
/*
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		} else {
			counts[input.Text()]++
		}
	}

	for text, number := range counts {
		if number > 1 {
			fmt.Printf("%d\t%s\n", number, text)
		}
	}
}*/

// 从文件中读取重复的行 (以"流"模式读取输入，并根据需要拆分成多个行)
/*func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "version2 : %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for filename, textAndNumber := range counts {
		for text, number := range textAndNumber {
			if number > 1 {
				fmt.Println("filename: ", filename, " text: ", text, " number: ", number)
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	if filename == "" {
		for input.Scan() {
			counts["stdin"][input.Text()]++
		}
	} else {
		counts[filename] = make(map[string]int)
		for input.Scan() {
			counts[filename][input.Text()]++
		}
	}
}*/

// 使用ReadFile函数读取文件，(将整个文件读取到内存中)
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) // ReadFile函数返回一个字节切片，必须把它转换成string，才能使用strings.Split进行分割
		if err != nil {
			fmt.Fprintf(os.Stderr, "version3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("text: ", line, " number: ", n)
		}
	}
}
