package slc

func FilterInt(data []int, f func(int) bool) []int {
	newData := []int{}

	for i := range data {
		
		if (f(data[i])) {
			newData = append(newData, data[i])
		}

	}

	return newData
}

func FilterFloat(data []float64, f func(float64) bool) []float64 {
	newData := []float64{}

	for i := range data {
		
		if (f(data[i])) {
			newData = append(newData, data[i])
		}
		
	}

	return newData
}