package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	go responseSize("https://www.duoke360.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(time.Second * 10)
	return
}
