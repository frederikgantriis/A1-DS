package main

import "fmt"

func main() {
	var channels []chan bool
	for i := 0; i < 10; i++ {
		channels = append(channels, make(chan bool))
	}

	for i := 0; i < 5; i++ {

		// the calculations inside the brackets ensures that the correct channels
		// are given to the correct philospher/fork eg: philosopher 0 gets
		// fork 0*2=0 and 0*2+1=1
		go philosopher(channels[i*2], channels[i*2+1])
		go chopstick(channels[i*2+1], channels[i*2+2%10])
	}
}

func philosopher(c1 chan bool, c2 chan bool) {
	meals := 0

	for meals < 2 {
		// TODO: Implement philosopher function
		result1 := <-c1

		if result1 == true {
			// TODO: Implement waiting timer for channel2
			result2 := <-c2

			if result2 == true {
				fmt.Print("EATEN")
				meals++
				c1 <- true
				c2 <- true
			} else {
				fmt.Print("THINKING")
				c1 <- true
			}
		}

	}
}

func chopstick(c1 chan bool, c2 chan bool) {
	for {
		// TODO: Implement chopstick function
	}
}
