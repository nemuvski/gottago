package main

import "fmt"

const MAX_NUM_ITERATIONS = 10
const THRESHOLD_DIFF = 0.0001

func sqrt(x float64) float64 {
	z := 1.0
	pre_z := x

	for i := 0; i < MAX_NUM_ITERATIONS; i++ {
		z -= (z*z - x) / (2 * z)

		// 前回のz値との差分が一定値を下回る場合は計算を止める
		if (pre_z - z) < THRESHOLD_DIFF {
			break
		}

		pre_z = z
	}

	return z
}

func main() {
	x := 2
	result := sqrt(float64(x))
	fmt.Printf("sqrt(%v) ≒ %v\n", x, result)
}
