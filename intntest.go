package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	headCount := 0
	for i := 0; i < 1000000; i++ {
		if rand.Intn(2) == 1 {
			headCount++
		}
	}

	fmt.Printf("%d\n", headCount)
}
