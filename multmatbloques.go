package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX   = 2000
	BLOCK = 100
)

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

func multiplicarMatricesBloques() {
	for ii := 0; ii < MAX; ii += BLOCK {
		for jj := 0; jj < MAX; jj += BLOCK {
			for kk := 0; kk < MAX; kk += BLOCK {
				for i := ii; i < min(ii+BLOCK, MAX); i++ {
					for j := jj; j < min(jj+BLOCK, MAX); j++ {
						for k := kk; k < min(kk+BLOCK, MAX); k++ {
							C[i][j] += A[i][k] * B[k][j]
						}
					}
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	iniciar()
	start := time.Now()
	multiplicarMatricesBloques()
	fmt.Printf("Tiempo para tamano %d: %v\n", MAX, time.Since(start))
}
