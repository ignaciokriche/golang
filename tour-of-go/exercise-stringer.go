// Exercise: Stringers
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%d.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func StringerExercise() {

	fmt.Println("Exercise: Stringers")

	hosts := map[string]IPAddr{
		"example":   {1, 2, 3, 4},
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"empty":     {},
	}

	for name, ip := range hosts {
		fmt.Printf("\t%v: %v\n", name, ip)
	}

	fmt.Println("----------------------------------------------------")
}
