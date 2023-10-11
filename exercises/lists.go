// Author: Ignacio Krichevsky

package exercises

import "fmt"

// A list is a pointer to a ListNode struct:
type ListNode struct {
	element any
	next    *ListNode
}

// Removes duplicates from a linked list.
func (list *ListNode) RemoveDuplicates() {
	// a temporary buffer is not allowed for this solution:
	for ; list != nil; list = list.next {
		for previous, current := list, list.next; current != nil; current = current.next {
			if current.element == list.element {
				// skip current
				previous.next = current.next
			} else {
				previous = current
			}
		}
	}
}

func (list *ListNode) String() string {
	switch {
	case list == nil:
		return "<nil>"
	case list.next == nil:
		return fmt.Sprintf("%v", list.element)
	default:
		return fmt.Sprintf("%v, %v", list.element, list.next)
	}
}
