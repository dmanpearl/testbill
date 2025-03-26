package proc

import (
	"testing"
)

func TestMultOfGroupSize(t *testing.T) {
	results := Process([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	if len(results) != 3 {
		t.Errorf("group count: %d, want: %d\n", len(results), 3)
	}
	if !results[len(results)-1].Whole {
		t.Errorf("last group not whole\n")
	}
}

func TestBigGroupSize(t *testing.T) {
	results := Process([]int{1, 2}, 20)
	if len(results) != 1 {
		t.Errorf("group count: %d, want: %d\n", len(results), 1)
	}
	result := results[len(results)-1]
	if result.Whole {
		t.Errorf("group is incorrectly whole\n")
	}
	if result.Sum != 3 {
		t.Errorf("group sum: %d, want: %d\n", result.Sum, 3)
	}
}

func TestMultiples(t *testing.T) {
	for i, vals := range [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7},
		{8, 9, 10, 11, 12},
		{29, 31, 8, 17, 64, 92, 28, 1, 2, 3, 4, 6, 7, 9},
	} {
		groupSize := i + 1
		results := Process(vals, groupSize)
		wantGroups := (len(vals) + groupSize - 1) / groupSize
		if len(results) != wantGroups {
			t.Errorf("incorrect group count %d, want %d, i %d, vals %v\n", len(results), wantGroups, i, vals)
		}

		var wantSum int
		for j := 0; j < groupSize; j++ {
			wantSum += vals[j]
		}
		if results[0].Sum != wantSum {
			t.Errorf("incorrect first result.sum %d, want %d, i %d, vals %v\n", results[0].Sum, wantSum, i, vals)
		}
	}
}
