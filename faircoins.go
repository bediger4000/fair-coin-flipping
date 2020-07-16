package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
You have n fair coins and you flip them all at the same time. Any that come up
tails you set aside. The ones that come up heads you flip again. How many
rounds do you expect to play before only one coin remains?

Write a function that, given n, returns the number of rounds you'd expect to
play until one coin remains.
*/

const iterations float64 = 100000.0

func main() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// I think this is the function that satisfies the problem:
	fmt.Printf("%d fair coins, log2(%d) = %f\n", n, n, math.Log2(float64(n)))

	// Nevertheless, do iterations number of simulations to see
	// how close we came.
	sumCount := 0
	for i := 0; i < int(iterations); i++ {
		sumCount += flipCoins(n)
	}

	fmt.Printf("%.0f iterations gives %f average attempts\n", iterations, float64(sumCount)/iterations)

}

// flipCoins performs rounds of mass coin flipping until
// at most 1 coin remains that came up "tails".
func flipCoins(n int) int {
	var i int
	for i = 0; n > 1; i++ {
		n = round(n)
	}
	return i
}

// round "flips" n fair coins, and returns the number that comes up heads.
func round(n int) int {
	headCount := 0
	for i := 0; i < n; i++ {
		if rand.Intn(2) == 1 {
			headCount++
		}
	}
	return headCount
}
