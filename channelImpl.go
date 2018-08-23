package main

import (
	"fmt"
	"sync"
)

type Mmsgs struct {
	Message string
}

func ping(pings chan<- Mmsgs, msg string) {
	pings <- Mmsgs{msg}
}

func pong(pings <-chan Mmsgs, pongs chan<- Mmsgs) {
	msg := <-pings
	pongs <- msg
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	pings := make(chan Mmsgs)
	pongs := make(chan Mmsgs)
	pings <- Mmsgs{msg}
	go ping(pings, "passed message")
	go ping(pings, "passing again")
	go pong(pings, pongs)
	ret := <-pongs
	//close(pings)
	//close(pongs)
	fmt.Println(ret.Message)
}
