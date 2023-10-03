/**
 * You are given a string of digits such as "123456"
 * your goal is to write a method to sum each digit from left to right and
 * if the sum has more than 2 digits then you first need to sum those digits
 * until only one digit is left, and then you keep moving.
 *
 * For the example above: "123456"
 * 1 + 2 + 3 + 4 = 10
 * 1 + 0 = 1
 * 1 + 5 = 6
 * 6 + 6 = 12
 * 1 + 2 = 3
 * then sumDigits("123456") = 3
 *
 * For: "99"
 * 9 + 9 = 18
 * 1 + 8 = 9
 * then sumDigits("99") = 9
 *
 * For: "14522"
 * 1 + 4 + 5 = 10
 * 1 + 0 = 1
 * 1 + 2 = 3
 * 3 + 2 = 5
 *
 * then sumDigits("14522") = 5
 *
 * Author: Ignacio Krichevsky
 *
 */

package exercises

import "strconv"

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
