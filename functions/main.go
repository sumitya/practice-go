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

	values := [][]int{[]int{1, 2}, []int{3, 4}} // two dim. int slice
	printer(values)

	aSlice := []int{3, 2, 1, 0}
	rangeOverIntCollection(sortASlice(aSlice))

	printer(reverseAString("abcd"))
	printer(reverseString("name"))
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

func sortASlice(aSlice []int) []int {
	// pass through a slice
	for pass := 1; pass < len(aSlice); pass++ {
		// one pass
		for i := 0; i < len(aSlice)-pass; i++ {
			if aSlice[i] > aSlice[i+1] {
				aSlice[i], aSlice[i+1] = aSlice[i+1], aSlice[i]
			}
		}
	}
	return aSlice
}

func reverseAString(str string) string {
	rns := []rune(str) // this convert string to slice of rune. "name" -> ['n','a','m','e']
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func reverseString(str string) string {
	runs := []rune(str)

	for i, j := 0, len(runs)-1; i < j; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]
	}

	return string(runs)
}
