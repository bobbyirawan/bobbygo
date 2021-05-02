package bobbygo

func QuickSort(Numbers []float64, as string) []float64 {
	newArr := make([]float64, len(Numbers))

	for i, v := range Numbers {
		newArr[i] = v
	}

	sort(newArr, 0, len(Numbers)-1, as)

	return newArr
}

func sort(Numbers []float64, start, end int, as string) {
	if (end - start) < 1 {
		return
	}

	pivot := Numbers[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if as == "desc" {
			if Numbers[i] > pivot {
				temp := Numbers[splitIndex]
				Numbers[splitIndex] = Numbers[i]
				Numbers[i] = temp
				splitIndex++
			}
		} else {
			if Numbers[i] < pivot {
				temp := Numbers[splitIndex]
				Numbers[splitIndex] = Numbers[i]
				Numbers[i] = temp
				splitIndex++
			}
		}
	}

	Numbers[end] = Numbers[splitIndex]
	Numbers[splitIndex] = pivot

	sort(Numbers, start, splitIndex-1, as)
	sort(Numbers, splitIndex+1, end, as)
}
