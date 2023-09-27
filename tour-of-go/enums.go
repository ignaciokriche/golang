// Within a constant declaration, the predeclared identifier iota represents successive untyped integer
// constants. Its value is the index of the respective ConstSpec in that constant declaration, starting
// at zero. It can be used to construct a set of related constants:

package main

import "fmt"

func IotaDemo() {

	fmt.Println("enums in go:")

	const (
		ZERO = iota
		ONE
		TWO
		THREE
	)

	fmt.Println(ZERO)
	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE)
	fmt.Println()

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
		bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
		_, _                                 //                        (iota == 2, unused)
		bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
	)

	fmt.Println(bit0, mask0)
	fmt.Println(bit1, mask1)
	fmt.Println(bit3, mask3)
	fmt.Println()

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}
