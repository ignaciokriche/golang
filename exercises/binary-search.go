// Author: Ignacio Krichevsky

package exercises

// Returns the index of n in numbers or -1 if not present.
func BinarySearchRecursive(numbers []int, n int) int {

	if len(numbers) == 0 {
		return -1
	}

	i := len(numbers) / 2
	value := numbers[i]
	switch {
	case value < n:
		return BinarySearchRecursive(numbers[i+1:], n)
	case value == n:
		return i
	default:
		return BinarySearchRecursive(numbers[:i], n)
	}

}

// Returns the index of n in numbers or -1 if not present.
func BinarySearch(numbers []int, n int) int {

	for len(numbers) > 0 {

		i := len(numbers) / 2
		value := numbers[i]

		switch {
		case value < n:
			numbers = numbers[i+1:]
		case value == n:
			return i
		default:
			numbers = numbers[0:i]
		}
	}
	return -1
}
