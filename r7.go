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

// rand5 returns an int in the range 0 to 4 inclusive,
// using a prescribed function rand7, that gives back uniformly
// distributed integers in the rand 0 to 6 inclusive.
func rand5() int {
	n := rand7()
	for n > 4 {
		n = rand7()
	}
	return n
}

// rand7 is the function described as given in the problem statement.
func rand7() int {
	return rand.Intn(7)
}
