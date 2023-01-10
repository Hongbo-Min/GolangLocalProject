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
	// WriteFile()
	// OpenClose()
	// ReadOps()
	Write()
	WriteString()
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
	err := os.Rename("text.txt", "2.txt")
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

func OpenClose() {
	f, err := os.Open("text.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	f2, err2 := os.OpenFile("text.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err2 != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}
}

func ReadOps() {
	/* f, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	for {
		buf := make([]byte, 10)
		n, err2 := f.Read(buf)
		fmt.Println("n : ", n, string(buf))
		if err2 == io.EOF {
			break
		}
	} */

	/* f, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	buf := make([]byte, 3)
	n, err2 := f.ReadAt(buf, 3)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("n: %v\n", n)
	}
	fmt.Printf("string(buf): %v\n", string(buf)) */

	de, err := os.ReadDir("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func Write() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Write([]byte("hello"))
	f.Close()
}

func WriteString() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteString("hello min hong bo")
	f.Close()
}
