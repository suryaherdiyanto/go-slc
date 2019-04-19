package slc

import (
	"reflect"
	"sort"
)

func doSquential(i, val interface{}) int {

	args := reflect.ValueOf(i)
	val := reflect.ValueOf(val)

	switch val.Kind() {
		case reflect.Int:
			val = val.Int()

			for index := 0; index < args.Len(); index++ {

				if (args.Index(index).Int() == val) {
					return index
				}

			}
		case reflect.Float64:
			val = val.Float()

			for index := 0; index < args.Len(); index++ {

				if (args.Index(index).Float() == val) {
					return index
				}

			}
	}
	

	return -1

}

func doBinaryInt(i []int, val int) int {

	args := i
	sort.Ints(args)

	head := 0
	tail := len(i)
	mid := int((tail + head) / 2)

	for index := 0; index < len(args); index ++ {
		
		if (args[mid] > val) {
			tail = mid
		}else{
			head = mid
		}

		mid = int((tail + head) / 2)

		if ((head == tail-1) && (args[mid] == val)) {
			return mid
		}

	}

	return -1

}

func doBinaryFloat(i []float64, val float64) int {

	args := i
	sort.Float64s(args)

	head := 0
	tail := len(i)
	mid := int((tail + head) / 2)

	for index := 0; index < len(args); index ++ {
		
		if (args[mid] > val) {
			tail = mid
		}else{
			head = mid
		}

		mid = int((tail + head) / 2)

		if ((head == tail-1) && (args[mid] == val)) {
			return mid
		}

	}

	return -1

}

func FindInts(data []int, val int) int {
	
	if (len(data) <= 10) {
		return doSquential(data, val)
	}

	return doBinaryInt(data, val)

}

func FindFloats(data []float64, val float64) int {

	if (len(data) <= 10) {
		return doSquential(data, val)
	}

	return doBinaryFloat(data, val)

}