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

var vals = [5][5]int{
	{1, 2, 3, 4, 5},
	{6, 7, 1, 2, 3},
	{4, 5, 6, 7, 1},
	{2, 3, 4, 5, 6},
	{7, 0, 0, 0, 0},
}

func rand7() int {
	n := rand5()
	m := rand5()
	r := vals[m][n]
	for r == 0 {
		r = vals[rand5()][rand5()]
	}
	return r
}

// rand5 is the function described as given in the problem statement.
func rand5() int {
	return rand.Intn(5)
}
