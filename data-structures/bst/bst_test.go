package bst

import "testing"

func TestBst_Insert_Search(t *testing.T) {
	want := []int{5, 3, 8, 2, 4, 7, 9}
	bst := &Tree{}

	for _, v := range want {
		bst.Insert(v)
	}
	
	// positive assertion
	for _, v := range want {
		if !bst.Search(v) {
			t.Fatalf("expected %d but was not found in tree", v)
		}
	}

	// negative assertion
	for _, v := range []int{1, 10, 15, 40} {
		if bst.Search(v) {
			t.Fatalf("expected not to find %d in tree", v)
		}
	}
}