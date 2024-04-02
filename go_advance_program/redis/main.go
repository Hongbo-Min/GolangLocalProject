package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	lockKey := "counter_lock"
	counterKey := "counter"
	resp := client.SetNX(context.Background(), lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}
	getResp := client.Get(context.Background(), counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(context.Background(), counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			println("set value error!")
		}
		println("current counter is ", cntValue)
		delResp := client.Del(context.Background(), lockKey)
		unlockSuccess, err := delResp.Result()
		if err == nil && unlockSuccess > 0 {
			println("unlock success!")
		} else {
			println("unlock failed ", err)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
