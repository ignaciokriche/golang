// Author: Ignacio Krichevsky

package exercises

func Anagrams(str string) []string {

	if len(str) == 1 {
		return []string{str}
	}

	anagrams := []string{}

	for i := 0; i < len(str); i++ {

		subStr := str[:i] + str[i+1:]
		ch := str[i : i+1]

		for _, subA := range Anagrams(subStr) {
			anagrams = append(anagrams, ch+subA)
		}

	}

	return anagrams

}
