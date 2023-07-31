package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/rosedblabs/wal"
)

var KVMap map[string]*wal.ChunkPosition

func main() {
	KVMap = make(map[string]*wal.ChunkPosition)
	wal, _ := wal.Open(wal.DefaultOptions)

	for {
		err := userWriteData(wal)
		if err != nil {
			log.Println(err.Error())
		}

		err = userReadData(wal)
		if err != nil {
			log.Println(err.Error())
		}
	}

}

func userWriteData(wal *wal.WAL) error {
	var key string
	fmt.Println("请输入key: ")
	fmt.Scanln(&key)
	fmt.Println("输入的内容为: ", key)

	cp, err := wal.Write([]byte(key))
	if err != nil {
		return err
	}

	KVMap[key] = cp

	fmt.Println("write to wal success")

	return nil
}

func userReadData(wal *wal.WAL) error {
	var key string
	fmt.Println("请输入想要读取的key: ")
	fmt.Scanln(&key)
	fmt.Println("输入的内容为: ", key)

	if cp, ok := KVMap[key]; ok {
		b, err := wal.Read(cp)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	} else {
		return errors.New("key not exist")
	}

	return nil
}
