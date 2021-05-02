package bobbygo

func SelectionSort(numbers []float64, as string) []float64 {
	for i := 0; i < len(numbers)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if as == "desc" {
				if numbers[minIndex] < numbers[j] {
					minIndex = j
				}
			} else {
				if numbers[minIndex] > numbers[j] {
					minIndex = j
				}
			}
		}
		tmp := numbers[i]
		numbers[i] = numbers[minIndex]
		numbers[minIndex] = tmp
	}

	return numbers
}

