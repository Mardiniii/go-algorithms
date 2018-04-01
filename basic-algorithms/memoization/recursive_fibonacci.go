package fibonacci

// Time complexity: O(2^n)
// Space complexity: O(n)

// Execution time using this approach: 0.204s

// go test recursive_fibonacci_test.go recursive_fibonacci.go
// ok      command-line-arguments  0.204s

// RecursiveFibonacci calculates the sequence using recursion calls
func RecursiveFibonacci(n int) int {
	if n < 2 {
		return n
	}

	return RecursiveFibonacci(n-2) + RecursiveFibonacci(n-1)
}
