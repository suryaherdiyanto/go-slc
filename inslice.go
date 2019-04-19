package slc

func IsInSliceInts(i []int, val int) bool {

	if (FindInts(i, val) != -1) {
		return true
	}

	return false

}

func IsInSliceFloats(i []float64, val float64) bool {

	if (FindFloats(i, val) != -1) {
		return true
	}

	return false

}