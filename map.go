package slc

func MapInt(data []int, f func(int) int) []int {
	newData := []int{}

	for i := range data {
		newData = append(newData, f(data[i]))
	}

	return newData

}

func MapFloat(data []float64, f func(float64) float64) []float64 {
	newData := []float64{}
	
	for i := range data {
		newData = append(newData, f(data[i]))
	}
	
	return newData

}