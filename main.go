package main

func main() {
	var channels []chan bool
	for i := 0; i < 5; i++ {
		channels = append(channels, make(chan bool))
	}

	for i := 0; i < 5; i++ {
		go philosopher(channels[i], channels[i+1%5])
		go chopstick(channels[i])
	}
}

func philosopher(c1 chan bool, c2 chan bool) {
	for {
		// TODO: Implement philosopher function
	}
}

func chopstick(c chan bool) {
	for {
		// TODO: Implement chopstick function
	}
}
