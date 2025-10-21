package datastructures

import "testing"

func TestSet_Dupes(t *testing.T) {
	s := NewSet[string]()

	// Add values inc dupes
	s.Add("apple")
	s.Add("banana")
	s.Add("apple")

	// Remove "apple"
	s.Remove("apple")

	// "apple" is removed completely because dupe was igored
	if s.Has("apple") {
		t.Fatalf("expected record to have been removed")
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[int]()
	for _, v := range []int{1,2,3} {
		s1.Add(v)
	}

	s2 := NewSet[int]()
	for _, v := range []int{3,4,5} {
		s2.Add(v)
	}

	unionS := s1.Union(s2)

	// Expected 1, 2, 3, 4, 5
	want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true }

	if len(unionS.data) != len(want) {
		t.Fatalf("expected %d unique elements, got %d", len(want), len(unionS.data))
	}

	for k := range want {
		if !unionS.Has(k) {
			t.Fatalf("union missing element %d", k)
		}
	}
}

func TestSet_Union_TableDriven(t *testing.T) {
	tests := []struct{
		name string
		a, b []string
		want []string
	}{
		{ "disjoint", []string{"a", "b"}, []string{"c", "d"}, []string{"a", "b", "c", "d"}},
		{ "overlap", []string{"a", "b"}, []string{"b", "c"}, []string{"a", "b", "c"}},
		{ "subset", []string{"a", "b", "c"}, []string{"a"}, []string{"a", "b", "c"}},
		{ "emptyA", nil, []string{"x"}, []string{"x"}},
		{ "emptyB", []string{"y"}, nil, []string{"y"}},
		{ "bothEmpty", nil, nil, []string{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, b := NewSet[string](), NewSet[string]()
			for _, v := range tc.a {
				a.Add(v)
			}
			for _, v := range tc.b {
				b.Add(v)
			}
			
			got := a.Union(b)
			gotVals := got.Values()
			wantMap := make(map[string]bool, len(tc.want))
			for _, v := range tc.want {
				wantMap[v] = true
			}

			if len(gotVals) != len(wantMap) {
				t.Fatalf("got %d elements, want %d elements", len(gotVals), len(wantMap))
			}

			for _, v := range gotVals {
				if !wantMap[v] {
					t.Fatalf("unexpected element in union: %v", v)
				}
			}
		})
	}
}