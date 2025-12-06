package main

import "fmt"

func main() {
	var arrAge1 = [5]int{5, 6, 7, 8, 8}
	var arrAge2 [5]int
	arrSlice1 := make([]int, 5) // init slice with 5 size
	arrSlice2 := []int{2, 3, 4} // init slice
	var arrSlice3 []int

	ptr := new([]int)       // it declares a pointer, which is nil
	slice := make([]int, 0) // it declares a slice, which points to empty array

	rangeOverIntCollection(slice)
	printer(ptr)

	funcWithArrayParam(arrAge1)
	funcWithArrayParam(arrAge2)
	funcWithSliceParam(arrSlice1)
	funcWithSliceParamPtr(&arrSlice2)
	funcWithSliceParam(arrSlice3)

	printer("abcd")

}

func funcWithSliceParamPtr(arrAge *[]int) {
	rangeOverIntCollection(*arrAge)
}
func funcWithSliceParam(arrAge []int) {
	rangeOverIntCollection(arrAge)
}
func funcWithArrayParam(ageArr [5]int) {
	rangeOverIntCollection(ageArr[:]) // ageArr[:] , here slice is equal to orginial array
}

func rangeOverIntCollection(coll []int) {
	for _, item := range coll {
		fmt.Println(item)
	}
}

func printer(variable any) {
	fmt.Println(variable)
}
