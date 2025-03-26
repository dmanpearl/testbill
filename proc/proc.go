package proc

type result struct {
	Sum   int
	Whole bool
}

func Process(vals []int, groupSize int) []result {
	var results []result

	// Print all values in the input slice
	// for i, val := range vals {
	// 	fmt.Printf("%d. %v\n", i, val)
	// }

	// Print the sum of the next groupSize values in the input slice
	// (not the requirement)
	// for i := 0; i < len(vals)-groupSize; i++ {
	// 	sum := 0
	// 	for j := 0; j < groupSize; j++ {
	// 		sum += vals[i+j]
	// 	}
	// 	fmt.Printf("%d. %v\n", i, sum)
	// }

	// Print the sum of each group of groupSize values. indicate groups that could not satisfy the groupSize quantity.
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
		// status := "complete"
		// if !whole {
		// 	status = "partial"
		// }
		// fmt.Printf("%d. %v %v\n", group, sum, status)
		results = append(results, result{Sum: sum, Whole: whole})
	}
	return results
}
