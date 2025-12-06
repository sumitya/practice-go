package main

import "fmt"

func main() {
	var arrAge1 = [5]int{5, 6, 7, 8, 8}
	var arrAge2 [5]int
	arrAge3 := make([]int, 5) // init slice with 5 size
	arrAge4 := []int{2, 3, 4} // init slice with 5 size

	funcWithArrayParam(arrAge1)
	funcWithArrayParam(arrAge2)
	funcWithSliceParam(arrAge3)
	funcWithSliceParamPtr(&arrAge4)

}

func funcWithSliceParamPtr(arrAge *[]int) {
	rangeOverIntCollection(*arrAge)
}
func funcWithSliceParam(arrAge []int) {
	rangeOverIntCollection(arrAge)
}
func funcWithArrayParam(ageArr [5]int) {
	rangeOverIntCollection(ageArr[:])
}

func rangeOverIntCollection(coll []int) {
	for _, item := range coll {
		fmt.Println(item)
	}
}
