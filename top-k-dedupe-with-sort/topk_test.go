package topkdedupewithsort

import (
	"testing"
)

func TestTopKUsers_Smoke(t *testing.T) {
	users := []User{
		{ID: 1, Name: "A", Score: 72},
		{ID: 2, Name: "B", Score: 85},
		{ID: 3, Name: "C", Score: 60},
		{ID: 4, Name: "D", Score: 91},
		{ID: 5, Name: "E", Score: 88},
	}
	got := TopKUsers(users, 3)
	// Expect top 3 by Score: D(91), E(88), B(85)
	want := []User{{4, "D", 91}, {5, "E", 88}, {2, "B", 85}}

	if !equalUsers(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTopKUsers_EdgeCases(t *testing.T) {
	cases := []struct {
		name  string
		users []User
		k     int
		want  []User
	}{
		{"k<=0", []User{{1, "A", 10}}, 0, nil},
		{"empty", nil, 3, nil},
		{"k>=n", []User{{1, "A", 2}, {2, "B", 5}}, 5, []User{{2, "B", 5}, {1, "A", 2}}},
		{"ties_by_score_id", []User{{2, "B", 10}, {1, "A", 10}, {3, "C", 9}}, 2, []User{{1, "A", 10}, {2, "B", 10}}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := TopKUsers(tc.users, tc.k)
			if !equalUsers(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// equalUsers compares only (ID, Score) pairs for brevity (names may vary).
func equalUsers(a, b []User) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ID != b[i].ID || a[i].Score != b[i].Score {
			return false
		}
	}
	return true
}
