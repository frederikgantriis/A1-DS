package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Deadlock doesn't occur since the philosophers check if both forks are
// available before trying to eat and always puts the forks back after use

func main() {
	rand.Seed(time.Now().Unix())
	var channels []chan bool
	var finishedEating = make(chan int)
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
		fmt.Printf("Philosopher %d has eaten three times\n", <-finishedEating)
	}
}

func philosopher(i int, out1 chan bool, in1 chan bool, out2 chan bool, in2 chan bool, finishedEating chan int) {
	meals := 0

	for {
		hasFork1 := <-in1
		out1 <- false
		hasFork2 := <-in2
		out2 <- false

		if hasFork2 && hasFork1 {
			meals++
			fmt.Printf("Philosopher %d eating\n", i)
			<-in1
			out1 <- true
			<-in2
			out2 <- true
			time.Sleep(100 * time.Millisecond)
			if meals == 3 {
				finishedEating <- i
				break
			}
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
