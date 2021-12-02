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
		fmt.Printf("%d\n", rand7())
	}
}

func rand7() int {
	n := rand5()
	m := rand5()
	if n+m < 7 {
		n += m
	}
	return n
}

// rand5 is the function described as given in the problem statement.
func rand5() int {
	return rand.Intn(5)
}
