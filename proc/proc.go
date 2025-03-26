package proc

type result struct {
	Sum   int
	Whole bool
}

// Generate the sum of each group of groupSize values.
// Indicate groups that could not satisfy the groupSize quantity with 'whole = False'.
func Process(vals []int, groupSize int) []result {
	groups := (len(vals) + groupSize - 1) / groupSize
	results := make([]result, groups)
	sum := 0
	for i, val := range vals {
		groupIdx := i / groupSize
		if groupIdx*groupSize == i {
			sum = 0 // New group
		}
		sum += val
		results[groupIdx] = result{Sum: sum, Whole: i == groupSize*(groupIdx+1)-1}
	}
	return results
}
