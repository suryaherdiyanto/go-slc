package slc

func FindInt(data []int, val int) int {
	var index int

	for i := range data {

		if data[i] == val {
			index = i
			break
		}

		if i == len(data)-1 {
			index = -1
		}

	}

	return index
}

func FindFloat(data []float64, val float64) int {
	var index int

	for i := range data {

		if data[i] == val {
			index = i
			break
		}

		if i == len(data)-1 {
			index = -1
		}

	}

	return index
}