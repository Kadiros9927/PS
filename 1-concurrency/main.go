package main

import (
	"fmt"
	"math/rand/v2"
)

func randomInt(chRandom chan<- int) {
	defer close(chRandom)
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.IntN(101)
		chRandom <- arr[i]
	}
}

func power(chRandom <-chan int, chPow chan<- int) {
	defer close(chPow)
	for num := range chRandom {
		chPow <- num * num
	}
}

func main() {
	chRandom := make(chan int)
	chPower := make(chan int)

	go randomInt(chRandom)
	go power(chRandom, chPower)

	for num := range chPower {
		fmt.Println(num)
	}
}
