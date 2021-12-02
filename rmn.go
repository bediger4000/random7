package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var vals [][]int

func main() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	m, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	if m <= n {
		log.Fatalf("%d has to be greater than %d\n", m, n)
	}
	if m > n*n {
		log.Fatalf("%d has to be greater than %d*%d\n", m, n, n)
	}

	fmt.Fprintf(os.Stderr, "Should end up with about %d of each value\n", max/m)

	valuesSetup(m, n)

	for i := 0; i < max; i++ {
		fmt.Printf("%d\n", randmn(n))
	}
}

func randmn(N int) int {
	r := vals[rand.Intn(N)][rand.Intn(N)]
	for r == 0 {
		r = vals[rand.Intn(N)][rand.Intn(N)]
	}
	return r
}

func valuesSetup(m, n int) {

	for i := 0; i < m; i++ {
		vals = append(vals, make([]int, m))
	}

	d := (n * n / m) * m

	v := 0
	for i := 0; i < d; i++ {
		v++
		if v > m {
			v = 1
		}
		vals[i/n][i%n] = v
	}

}
