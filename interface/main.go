package main

import (
	"fmt"
	"sort"
)

func main() {
	persons := Persons{
		{"Alice", 30},
		{"Bob", 25},
		{"Alpha", 34},
	}
	printer(persons)
	sort.Sort(persons)
	printer(persons)

	sort.Sort(sort.Reverse(persons)) // reverse the sorted collection
	printer(persons)

}

// this is an existing interface in sort package. by implementing its methods. we can define custom sort, reverse etc.

//	type Interface interface {
//		Len() int
//		Less(i, j int) bool
//		Swap(i, j int)
//	}
func printer(variable any) {
	fmt.Println(variable)
}

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (persons Persons) Len() int {
	return len(persons)
}

func (persons Persons) Less(i, j int) bool {
	return persons[i].Age < persons[j].Age
}

func (persons Persons) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}
