// uses a map to get and put elements in O(1) time.
// uses a doubly linked list (as a queue) to keep track of the next element to be
// removed when cache is full.
//
// Author: Ignacio Krichevsky

package exercises

import "container/list"

type LRUCache struct {

	// a map to store and retrieve in O(1) time.
	// the key of the map is the user key.
	// the value stored in the map is the user value and a pointer to
	// the queue element holding the map key.
	nodeMap map[int]*node

	// doubly linked queue to know what elemet to remove, if need to.
	nodeQueue *list.List

	Capacity int
}

type node struct {
	// the user value
	value int

	// a pointer to the queue element, where the element is the user key
	keyPtr *list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		make(map[int]*node, capacity),
		list.New(),
		capacity,
	}
}

func (cache *LRUCache) Get(key int) int {

	if node, ok := cache.nodeMap[key]; ok {
		// move it to the tail of the queue:
		cache.nodeQueue.MoveToBack(node.keyPtr)
		return node.value
	}

	return -1

}

func (cache *LRUCache) Put(key, value int) {

	if node, ok := cache.nodeMap[key]; ok {
		node.value = value
		// move it to the tail of the queue:
		cache.nodeQueue.MoveToBack(node.keyPtr)
		return
	}

	if len(cache.nodeMap) == cache.Capacity {
		// remove LRU from queue and map:
		nodeToDelete := cache.nodeQueue.Front()
		cache.nodeQueue.Remove(nodeToDelete)
		delete(cache.nodeMap, nodeToDelete.Value.(int))
	}

	// add new elem to queue and map:
	pointerToElem := cache.nodeQueue.PushBack(key)
	cache.nodeMap[key] = &node{value, pointerToElem}

}
