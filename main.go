package main

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
	for {
		// TODO: Implement philosopher function

	}
}

func chopstick(c1 chan bool, c2 chan bool) {
	for {
		// TODO: Implement chopstick function
	}
}
