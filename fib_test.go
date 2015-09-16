package fib

import "testing"
import "fmt"

func check(t *testing.T, v int, expected int) {
	r := calc_fib(v)
	if r != expected {
		t.Error("Expected ", expected, " but got ", r)
	}
}

func TestFib(t *testing.T) {
	check(t, 0, 0)
	check(t, 1, 1)
	check(t, 2, 1)
	check(t, 3, 2)
	check(t, 4, 3)
	check(t, 5, 5)
	check(t, 10, 55)

	fmt.Println(calc_fib(20))
}
