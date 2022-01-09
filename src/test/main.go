package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	sendData3(ch)
	go getDate3(ch)
}

func sendData3(ch chan string) {
	ch <- "ws"
	ch <- "cq"
	ch <- "bj"
	close(ch)
}

func getDate3(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Println(input)
	}
}
