package fibonacci

// Input: x - Value to use in a function
// Output: f(x) - Function result after evaluate the given value
// How it works: Memoization is a technique for improving the
// performance of recursive algorithms. It involves rewriting the
// recursive algorithm so that as answers to problems are found,
// they are stored in an data structure for example like an array.
// Recursive calls can look up results in the array rather than having
// to recalculate them. Memoization improves performance because
// partial results are never calculated twice
// Time complexity: O(n)
// Space complexity: O(n)

// Execution time using this approach: 0.008

// go test memoized_fibonacci_test.go memoized_fibonacci.go
// ok      command-line-arguments  0.008s

// This approach is almost 25x faster than the recursive solution

// MemoizedFibonacci calculates the sequence using memoization
// technique reducing time and space complexity
func MemoizedFibonacci(n int) int {
	results := []int{0, 1}

	for i := 2; i <= n; i++ {
		results = append(results, results[i-2]+results[i-1])
	}

	return results[n]
}
