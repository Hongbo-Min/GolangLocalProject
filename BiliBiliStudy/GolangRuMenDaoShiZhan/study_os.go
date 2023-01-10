package main

import (
	"fmt"
	"os"
)

func main() {
	// CreateFile("min.txt")
	// CreateDir("min")
	// Remove()
	// GetPWD()
	// Rename()
	// ReadFile()
	WriteFile()
}

// https://pkg.go.dev/std 标准库的文档

// 创建文件
func CreateFile(name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func CreateDir(name string) {
	/* err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
	} */
	err := os.MkdirAll("min/hong/bo", os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// 删除文件
func Remove() {
	err := os.Remove("min.txt")
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// 删除目录
func RemoveAll() {
	err := os.RemoveAll("min")
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func GetPWD() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}

	os.Chdir("/home/minhongbo")
	dir2, err2 := os.Getwd()
	if err2 != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("dir2: %v\n", dir2)
	}

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func Rename() {
	err := os.Rename("1.txt", "2.txt")
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func ReadFile() {
	b, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("string(b[:]): %v\n", string(b[:]))
	}
}

func WriteFile() {
	err := os.WriteFile("test.txt", []byte("hong bo"), os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
