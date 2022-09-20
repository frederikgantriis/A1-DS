package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var channels []chan bool
	var finishedEating = make(chan bool)
	for i := 0; i < 10; i++ {
		channels = append(channels, make(chan bool, 1))
	}

	for i := 0; i < 5; i++ {

		// the calculations inside the brackets ensures that the correct channels
		// are given to the correct philospher/fork eg: philosopher 0 gets
		// fork 0*2=0 and 0*2+1=1
		go fork(channels[i*2], channels[i*2+1])
	}
	for i := 0; i < 5; i++ {
		go philosopher(i, channels[i*2], channels[i*2+1], channels[(i*2+2)%10], channels[(i*2+3)%10], finishedEating)
	}

	for i := 0; i < 5; i++ {
		<-finishedEating
		fmt.Println("finished eating")
	}
}

func philosopher(i int, out1 chan bool, in1 chan bool, out2 chan bool, in2 chan bool, finishedEating chan bool) {
	meals := 0

	for {
		hasFork1 := <-in1
		out1 <- false
		hasFork2 := <-in2
		out2 <- false

		if hasFork2 && hasFork1 {
			meals++
			if meals == 3 {
				finishedEating <- true
			}
			fmt.Printf("Philosopher %d eating\n", i)
			<-in1
			out1 <- true
			<-in2
			out2 <- true
			time.Sleep(100 * time.Millisecond)
		} else {
			if hasFork1 {
				<-in1
				out1 <- true
			}
			if hasFork2 {
				<-in2
				out2 <- true
			}
			fmt.Printf("Philosopher %d thinking\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func fork(in <-chan bool, out chan<- bool) {
	out <- true
	for {
		temp := <-in
		out <- temp
	}
}
