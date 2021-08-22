package slowfastpointer

import "fmt"

// MEDIUM

func FindHappyNumber(num int) bool {
	slow := squareSumOfDigits(num)
	fast := squareSumOfDigits(squareSumOfDigits(num))

	for fast != slow {
		slow = squareSumOfDigits(slow)
		fast = squareSumOfDigits(squareSumOfDigits(fast))
	}

	return fast == 1
}

func squareSumOfDigits(num int) int {
	digits := fmt.Sprint(num)
	var sum int
	for _, d := range digits {
		d = d - '0'
		sum += (int(d) * int(d))
	}

	return sum
}
