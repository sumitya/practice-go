package main

import (
	"fmt"
	"time"
)

func mainBkp() {
	go printer()

	go func(a int) {
		fmt.Printf("its me! %v", a)
	}(10)

	time.Sleep(time.Second)
}

func printer() {
	fmt.Println("its me!")
}
