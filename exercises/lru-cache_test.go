// Author: Ignacio Krichevsky

package exercises

import "testing"

func TestDiscardsLRU(t *testing.T) {

	var theTested *LRUCache = NewLRUCache(3)
	theTested.Put(1, 1)
	theTested.Put(2, 2)
	theTested.Put(3, 3)
	theTested.Put(4, 4)
	theTested.Put(5, 5)

	if theTested.Get(1) != -1 {
		t.Fatal("expecting -1")
	}

	if theTested.Get(2) != -1 {
		t.Fatal("expecting -1")
	}
	if theTested.Get(3) != 3 {
		t.Fatal("expecting 3")
	}

	if theTested.Get(4) != 4 {
		t.Fatal("expecting 4")
	}

	if theTested.Get(5) != 5 {
		t.Fatal("expecting 5")
	}

}

func TestGetRefreshKey(t *testing.T) {

	var theTested *LRUCache = NewLRUCache(5)
	theTested.Put(1, 1)
	theTested.Put(2, 2)
	theTested.Put(3, 3)
	theTested.Put(4, 4)
	theTested.Put(5, 5)

	if theTested.Get(1) != 1 {
		t.Fatal("expecting 1")
	}

	theTested.Put(6, 6)

	if theTested.Get(1) != 1 {
		t.Fatal("expecting 1")
	}

	if theTested.Get(2) != -1 {
		t.Fatal("expecting -1")
	}

	if theTested.Get(3) != 3 {
		t.Fatal("expecting 3")
	}

	if theTested.Get(4) != 4 {
		t.Fatal("expecting 4")
	}

	if theTested.Get(5) != 5 {
		t.Fatal("expecting 5")
	}
	if theTested.Get(6) != 6 {
		t.Fatal("expecting 6")
	}

}

func TestPutRefreshKey(t *testing.T) {

	var theTested *LRUCache = NewLRUCache(5)
	theTested.Put(1, 1)
	theTested.Put(2, 2)
	theTested.Put(3, 3)
	theTested.Put(4, 4)
	theTested.Put(5, 5)
	theTested.Put(1, 1)
	theTested.Put(6, 6)

	if theTested.Get(1) != 1 {
		t.Fatal("expecting 1")
	}

	if theTested.Get(2) != -1 {
		t.Fatal("expecting -1")
	}

	if theTested.Get(3) != 3 {
		t.Fatal("expecting 3")
	}

	if theTested.Get(4) != 4 {
		t.Fatal("expecting 4")
	}

	if theTested.Get(5) != 5 {
		t.Fatal("expecting 5")
	}

	if theTested.Get(6) != 6 {
		t.Fatal("expecting 6")
	}

}
