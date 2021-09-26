package epi

// 5.9 Enumerate All Primes to N

/* SOLUTION: Time O( N * LOG(LOG(N)) ) | Space O(N)
- initialize a boolean array of size n, that will serve as a prime number lookup table for numbers 1 to n
	and set all values > 1 to true
- for each number between 1 to n:
	- if number is prime:
		- add number to result
		- mark all the multiples of that number between number^2 to n as not prime
*/

func EnumeratePrimes(n int) []int {
	result := []int{}
	if n < 2 {
		return result
	}

	isPrime := make([]bool, n)
	for i := 1; i < n; i++ {
		isPrime[i] = true
	}

	for i := 1; i < n; i++ {
		if isPrime[i] {
			result = append(result, i+1)
			for j := i + 1; j*(i+1) <= n; j++ {
				isPrime[(j*(i+1))-1] = false
			}
		}
	}

	return result
}
