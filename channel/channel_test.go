package main

import (
	"testing"
	"fmt"
	"time"
)

func TestChannel(t *testing.T) {
	t.Run("testChannel", func(t *testing.T) {
		cha1 := make(chan int)
		go sendData(cha1)
		fmt.Println(<-cha1)
	})

	t.Run("testClose", func(t *testing.T) {
		ch := make(chan int)
		go producer(ch)
		for {
			v, ok := <-ch
			if ok == false {
				break
			}
			fmt.Println("Received ", v, ok)
		}
	})

	t.Run("test for range", func(t *testing.T) {
		ch := make(chan int)
		go producer(ch)
		for v := range ch {
			fmt.Println("Received ", v)
		}
	})

	t.Run("test buffer chan", func(t *testing.T) {
		ch := make(chan int, 2)
		go write(ch)
		time.Sleep(2 * time.Second)
		for v := range ch {
			fmt.Println("read value", v, "from ch")
			time.Sleep(2 * time.Second)
		}
	})
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}