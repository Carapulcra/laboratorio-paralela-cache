package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX = 10000

var A [MAX][MAX]float64
var x, y [MAX]float64

func iniciar() {
	rand.Seed(42)
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			A[i][j] = float64(rand.Intn(100))
		}
		x[i] = float64(i)
		y[i] = 0
	}
}

func first_pair() {
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			y[i] += A[i][j] * x[j]
		}
	}
}

func second_pair() {
	for j := 0; j < MAX; j++ {
		for i := 0; i < MAX; i++ {
			y[i] += A[i][j] * x[j]
		}
	}
}

func main() {
	iniciar()
	start := time.Now()
	first_pair()
	fmt.Printf("Tiempo first pair: %v\n", time.Since(start))

	iniciar()
	start = time.Now()
	second_pair()
	fmt.Printf("Tiempo second pair: %v\n", time.Since(start))
}
