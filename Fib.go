package fib

func calc_fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return calc_fib(n - 1) + calc_fib(n - 2)
}
