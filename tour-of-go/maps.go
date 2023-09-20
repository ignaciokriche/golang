package main

import "fmt"

func MapsDemo() {

	fmt.Println("maps demo")

	aMap := make(map[string]string)

	aMap["claudio"] = "krichevsky"
	aMap["edith"] = "rodriguez"
	aMap["moli"] = "canoli"
	aMap["ignacio"] = "krichevsky"

	for k, v := range aMap {
		fmt.Printf("map[%v]=%v\n", k, v)
	}

	checkAndPrint(aMap, "claudio")
	checkAndPrint(aMap, "salomon")

	aMap["salomon"] = "krichevsky"
	delete(aMap, "claudio")
	checkAndPrint(aMap, "claudio")
	checkAndPrint(aMap, "salomon")

	fmt.Println()

}

func checkAndPrint(aMap map[string]string, key string) {
	if elem, ok := aMap[key]; ok {
		fmt.Printf("%v key is present and its value is %v.\n", key, elem)
	} else {
		fmt.Printf("%v key is not present.\n", key)
	}
}
