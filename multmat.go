package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX = 2000

var A, B, C [MAX][MAX]float64

func iniciar() {
	rand.Seed(42)
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			A[i][j] = float64(rand.Intn(100))
			B[i][j] = float64(rand.Intn(100))
			C[i][j] = 0
		}
	}
}

func multiplicarMatrices() {
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			for k := 0; k < MAX; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func main() {
	iniciar()
	start := time.Now()
	multiplicarMatrices()
	fmt.Printf("Tiempo para tamano %d: %v\n", MAX, time.Since(start))
}
