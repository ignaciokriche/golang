// Reverts a string or a int.
//
// Author: Ignacio Krichevsky

package exercises

func Revert(v any) any {

	switch typedV := v.(type) {
	case string:
		return revertText(typedV)
	case int:
		return revertNumber(typedV)
	default:
		panic("don´t know how to revert.")
	}

}

func revertText(text string) string {
	if len(text) <= 1 {
		return text
	}
	return revertText(text[1:]) + text[0:1]
}

func revertNumber(number int) int {

	if number < 0 {
		return -revertNumber(-number)
	}

	reverted := 0
	for number > 0 {
		digit := number % 10
		number /= 10
		reverted = reverted*10 + digit
	}

	return reverted

}
