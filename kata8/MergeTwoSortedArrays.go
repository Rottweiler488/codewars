package kata8

import "sort"

func MergeArrays(arr1, arr2 []int) []int {
	merged := append(arr1, arr2...)

	sort.Ints(merged)

	var uniqueArray []int
	if len(merged) > 0 {
		uniqueArray = append(uniqueArray, merged[0])
	}
	for i := 1; i < len(merged); i++ {
		if merged[i] != merged[i-1] {
			uniqueArray = append(uniqueArray, merged[i])
		}
	}

	return uniqueArray
}
