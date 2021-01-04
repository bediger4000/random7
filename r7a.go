package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	for i := 0; i < 500000; i++ {
		fmt.Printf("%d\n", rand5())
	}
}

// rand5 returns the count of the number of times func rand7
// got invoked to get a number between 0-4 inclusive.
// This allows investigation of the probability of long run times
// of rand5
func rand5() int {
	n := rand7()
	i := 1
	for n > 4 {
		i++
		n = rand7()
	}
	return i
}

// rand7 is the function described as given in the problem statement.
func rand7() int {
	return rand.Intn(7)
}
