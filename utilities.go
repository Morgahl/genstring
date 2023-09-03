package genstring

import "golang.org/x/exp/constraints"

// assumed positve integer exponent
func intPow[N constraints.Integer](n, m N) N {
	if m == 0 {
		return 1
	}
	result := n
	for i := N(2); i <= m; i++ {
		result *= n
	}
	return result
}

func intLog2[N, M constraints.Integer](n N) (m M) {
	if n <= 0 {
		panic("log2Floor: n must be positive")
	}
	for ; n > 1; m++ {
		n >>= 1
	}
	return m
}
