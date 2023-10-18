// Author: Ignacio Krichevsky

package exercises

import "fmt"

func PrintHistogram(histo []int) {

	maxHeight := 0
	for _, height := range histo {
		if maxHeight < height {
			maxHeight = height
		}
	}

	for currentHeight := maxHeight; currentHeight > 0; currentHeight-- {
		for _, height := range histo {
			if height >= currentHeight {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

}
