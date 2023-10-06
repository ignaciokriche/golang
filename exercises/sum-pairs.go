/*
 * Given an integer array `arr` and an integer `k`,
 * find the pairs of elements in the array whose sum is k.
 *
 * Examples:
 *      arr = [1, 3, 4, 6, 3, 7, 10, 0, 9, -2]
 *      k = 7
 *      response: [(3, 4), (1, 6), (7, 0), (9, -2)]
 *
 *      arr = [1, 3, 4, 3, 4, 6, 3, 7, 10, 0, 0, 7, 7]
 *      k = 7
 *      response: [(3, 4), (3, 4), (1, 6), (7, 0), (0, 7)]
 *
 *
 * Author: Ignacio Krichevsky
 */

package exercises

func SumPairs(numbers []int, k int) []Pair {

	pairs := make([]Pair, 0, len(numbers)/2)
	complements := make(map[int]bool, len(numbers)/2)

	for _, n := range numbers {
		diff := k - n
		if complements[diff] {
			pairs = append(pairs, Pair{diff, n})
			complements[diff] = false
		} else {
			complements[n] = true
		}

	}

	return pairs

}

type Pair struct {
	X int
	Y int
}
