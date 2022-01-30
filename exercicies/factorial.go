package exercicies

func FactorialIteration(n int32) int32 {
	var result int32 = 1
	for i := int32(1); i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialRecusive(n int32) int32 {
	if n == 0 {
		return 1
	}
	return (n * FactorialRecusive(n-1))
}
