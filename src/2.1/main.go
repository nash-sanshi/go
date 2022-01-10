package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fmeng {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Printf("producer, %s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer, %s\n", data)
	}
	wg.Done()
}

func main() {
	//多个生产者和多个消费者模式
	chanSteam := make(chan string, 10)
	//生产者和消费者计数器
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)
	//3个生产者
	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanSteam)
	}
	//2个消费者
	for j := 0; j < 2; j++ {
		wgCs.Add(1)
		go consumer(wgCs, chanSteam)
	}

	//制造超时时间
	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()
	wgPd.Wait()
	//生产完成，关闭channel
	close(chanSteam)
	wgCs.Wait()
}
