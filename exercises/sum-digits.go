package exercises

import (
	"strconv"
)

func SumDigits(digits string) int {

	if digits == "" {
		return 0
	}

	if len(digits) == 1 {
		return int(digits[0] - '0')
	}

	// sum the first 2 digits which cannot be bigger than 9+9 = 18
	sum := int(digits[0] + digits[1] - 2*'0')

	// call recursivelly replacing the first 2 digits with its sum:
	return SumDigits(strconv.Itoa(sum) + digits[2:])

}

func SumDigitsIterative(digits string) int {
	sum := 0
	for _, ch := range digits {
		sum += int(ch - '0')
		if sum > 9 {
			sum = sum - 9
		}
	}
	return sum
}
