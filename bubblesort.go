package bobbygo

func BubbleSort(numbers []float64, as string) []float64 {

	for i := len(numbers); i > 0; i-- {
		if as == "desc" {
			for j := 1; j < i; j++ {
				if numbers[j-1] < numbers[j] {
					temp := numbers[j]
					numbers[j] = numbers[j-1]
					numbers[j-1] = temp
				}

			}
		} else {
			for j := 1; j < i; j++ {
				if numbers[j-1] > numbers[j] {
					temp := numbers[j]
					numbers[j] = numbers[j-1]
					numbers[j-1] = temp
				}
			}
		}
	}
	return numbers
}
