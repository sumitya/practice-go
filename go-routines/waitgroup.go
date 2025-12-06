package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ProducerConsumer()
}

func ProducerConsumer() {

	var waitGroup sync.WaitGroup

	for {
		randInt := rand.Int()
		fmt.Printf("produced!! %v \n", randInt)
		waitGroup.Add(1)
		go func(randInt int) {
			defer waitGroup.Done()
			fmt.Printf("consumed!! %v \n", randInt)
		}(randInt)

		time.Sleep(time.Second)
		waitGroup.Wait()
	}
}

func WaitGroupDemo() {
	var waitGroup sync.WaitGroup

	for i := 1; i <= 5; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
}
