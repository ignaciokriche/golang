package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func ExerciseEquivalentBinaryTreeDemo() {

	fmt.Println("Exercise Equivalent Binary Tree")

	tree1 := tree.New(100)
	tree2 := tree.New(100)

	fmt.Println("\ttree1:\t", tree1)
	fmt.Println("\ttree2:\t", tree2)

	nodes1Ch := make(chan int)
	go walkTreeAndClose(tree1, nodes1Ch)

	nodes2Ch := make(chan int)
	go walkTreeAndClose(tree2, nodes2Ch)

	if sameValuesAndOrder(nodes1Ch, nodes2Ch) {
		fmt.Println("\tBoth trees have the same values in the same order.")
	} else {
		panic("expecting both trees with same values in the same order.")
	}
	fmt.Println()

	fmt.Println("----------------------------------------------------")
	fmt.Println()
}

func walkTreeAndClose(t *tree.Tree, outCh chan int) {
	walkTree(t, outCh)
	close(outCh)
}

func walkTree(t *tree.Tree, outCh chan int) {

	if t == nil {
		return
	}

	walkTree(t.Left, outCh)
	outCh <- t.Value
	walkTree(t.Right, outCh)

}

func sameValuesAndOrder(ch1, ch2 chan int) bool {

	for {

		node1, opened1 := <-ch1
		node2, opened2 := <-ch2

		if !opened1 && !opened2 {
			// end of both channels at the same time.
			return true
		}

		if !opened1 || !opened2 {
			// one channel still has values and the other don´t.
			return false
		}

		if node1 != node2 {
			// values differ.
			return false
		}

	}

}
