package kata8

func MergeArrays(arr1, arr2 []int) []int {
	array := append(arr1, arr2...)
	ready := 0
	for i := 0; i < len(array); i++ {
		if i == len(array)-1 {
			ready = 0
			i = 0
		}

		if array[i] == array[i+1] {
			array = append(array[:i+1], array[i+2:]...)

			continue
		}

		if array[i] > array[i+1] {
			lastElement := array[i]
			array[i] = array[i+1]
			array[i+1] = lastElement

			continue
		}

		ready += 1

		if ready == len(array)-1 {
			break
		}
	}

	return array
}
