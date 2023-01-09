package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"
)

var (
	timeout      int64
	size         int
	count        int
	tp           uint8 = 8
	code         uint8 = 0
	sendCount    int
	successCount int
	failCount    int
	minTime      int64 = math.MaxInt32
	maxTime      int64
	totalTime    int64
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func getArgFromCommand() { // 通过flag包来解析命令行参数
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，单位毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位字节")
	flag.IntVar(&count, "n", 4, "发送请求次数")
	flag.Parse()
}

func main() {
	getArgFromCommand()
	desIp := os.Args[len(os.Args)-1] // 最后一个参数为ping操作目标地址
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	fmt.Printf("正在Ping %s [%s] 具有 %d 字节的数据\n", desIp, conn.RemoteAddr(), size)
	for i := 0; i < count; i++ {
		sendCount++
		tstart := time.Now()
		var icmp *ICMP = &ICMP{
			Type:        tp,
			Code:        code,
			CheckSum:    0,
			ID:          1,
			SequenceNum: 1,
		}

		data := make([]byte, size)
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		buffer.Write(data)
		data = buffer.Bytes()
		checksum := checkSum(data)

		data[2] = byte(checksum >> 8)
		data[3] = byte(checksum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		n, err := conn.Write(data)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}
		recive := make([]byte, 65535)
		n, err = conn.Read(recive)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}
		successCount++
		tsub := time.Since(tstart).Milliseconds()
		if tsub > maxTime {
			maxTime = tsub
		} else if tsub < minTime {
			minTime = tsub
		}
		totalTime += tsub
		fmt.Printf("正在加收来自 %d.%d.%d.%d 的回复：字节=%d 时间=%d ms TTL=%d\n", recive[12], recive[13], recive[14], recive[15], n-28, tsub, recive[8])
		time.Sleep(time.Second)
	}
	fmt.Printf("%s 的Ping统计信息：\n    数据包：已发送 = %d，已接受 = %d, 丢失 = %d（%.2f%%丢失）\n往返行程的估计时间(以毫秒为单位):\n 最短=%d ms，最长=%d ms\n",
		conn.RemoteAddr(), sendCount, successCount, failCount, float64(failCount)/float64(sendCount), minTime, maxTime)
}

// 计算校验和
func checkSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length != 0 {
		sum += uint32(data[index])
	}

	heigh16 := sum >> 16

	for heigh16 != 0 {
		sum = heigh16 + uint32(uint16(sum))
		heigh16 = sum >> 16
	}

	return uint16(^sum)
}
