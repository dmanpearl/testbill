package proc

type result struct {
	Sum   int
	Whole bool
}

func Process(vals []int, groupSize int) []result {
	var results []result
	// Generate the sum of each group of groupSize values. indicate groups that could not satisfy the groupSize quantity.
	groups := (len(vals) + groupSize - 1) / groupSize
	for group := 0; group < groups; group++ {
		i := group * 3
		sum := 0
		whole := true
		for j := 0; j < groupSize; j++ {
			if i+j >= len(vals) {
				whole = false
				break
			}
			sum += vals[i+j]
		}
		results = append(results, result{Sum: sum, Whole: whole})
	}
	return results
}
