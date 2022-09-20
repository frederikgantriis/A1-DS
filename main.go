package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var channels []chan bool
	for i := 0; i < 10; i++ {
		channels = append(channels, make(chan bool))
	}

	for i := 0; i < 5; i++ {

		// the calculations inside the brackets ensures that the correct channels
		// are given to the correct philospher/fork eg: philosopher 0 gets
		// fork 0*2=0 and 0*2+1=1
		go fork(channels[i*2], channels[i*2+1])
	}
	for i := 0; i < 5; i++ {
		go philosopher(i, channels[i*2], channels[i*2+1], channels[(i*2+2)%10], channels[(i*2+3)%10])
	}

	time.Sleep(10000 * time.Millisecond)
}

func philosopher(i int, out1 chan<- bool, in1 <-chan bool, out2 chan<- bool, in2 <-chan bool) {
	meals := 0

	for {
		if meals == 3 {

			fmt.Printf("Philosopher %d Finished\n", i)
			break
		}
		hasFork1 := <-in1
		if hasFork1 {
			out1 <- false
			fmt.Printf("Philosopher %d has Fork %d\n", i, i)
			hasFork2 := <-in2

			if hasFork2 {
				out2 <- false
				fmt.Printf("Philosopher %d has Fork %d\n", i, i+1)
				meals++
				n := rand.Intn(100)
				time.Sleep(time.Duration(n) * time.Millisecond)
				fmt.Printf("Philosopher %d eating\n", i)
				out1 <- true
				fmt.Printf("Philosopher %d dropped Fork %d\n", i, i)
				out2 <- true
				fmt.Printf("Philosopher %d dropped Fork %d\n", i, i+1)
			} else {
				fmt.Printf("Philosopher %d dropped Fork %d\n", i, i)
				out1 <- true
				fmt.Printf("Philosopher %d thinking\n", i)
				n := rand.Intn(100)
				time.Sleep(time.Duration(n) * time.Millisecond)
			}
		} else {
			fmt.Printf("Philosopher %d thinking\n", i)

			n := rand.Intn(100)
			time.Sleep(time.Duration(n) * time.Millisecond)
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
