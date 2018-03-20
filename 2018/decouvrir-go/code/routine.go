package main

import (
	"fmt"
	"time"
)

func worker() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Task #%d done.\n", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go worker()

	fmt.Println("Done.")

	time.Sleep(time.Second * 4)
}
