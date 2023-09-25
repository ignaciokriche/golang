package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func SimpleTreeExercisesDemo() {

	fmt.Println("Simple Tree Exercises:")
	fmt.Println()

	tree1 := tree.New(1)
	tree2 := tree.New(2)
	fmt.Println("tree1:", tree1)
	fmt.Println("tree2:", tree2)

	if equals(tree1, tree2) {
		fmt.Println("trees have the same values.")
	} else {
		fmt.Println("trees have not the same values.")
	}
	fmt.Println()

	tree3 := tree.New(3)
	tree3Copy := shallowCopy(tree3)
	fmt.Println("tree3:     ", tree3)
	fmt.Println("tree3 copy:", tree3Copy)
	fmt.Println()

	if equals(tree3, tree3Copy) {
		fmt.Println("trees have the same values.")
	} else {
		panic("expecting trees to be the same.")
	}
	fmt.Println()

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

func equals(t1, t2 *tree.Tree) bool {

	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Value == t2.Value &&
		equals(t1.Left, t2.Left) &&
		equals(t1.Right, t2.Right)
}

func shallowCopy(t *tree.Tree) *tree.Tree {

	if t == nil {
		return nil
	}

	return &tree.Tree{
		Left:  shallowCopy(t.Left),
		Value: t.Value,
		Right: shallowCopy(t.Right),
	}
}
