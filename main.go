package main

import (
	"errors"
	"fmt"
	fm "fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//prints
	fmt.Println("Hello, Go Project!")
	fmt.Println("hi duniya")
	fmt.Printf("%s", "dun")
	fmt.Println()

	fmt.Println(os.Getpid())
	fm.Println("another print")
	fm.Println(sum(1, 2))

	// data types
	var decision bool = true
	fmt.Println(decision)
	var ch byte = 'A'
	fmt.Println(ch)

	// logical operators
	fmt.Println(10 > 5)
	fmt.Println(10 > 5 && 10 < 5)
	fmt.Println(10 > 5 || 10 < 5)
	fmt.Println(!(10 > 5))
	fmt.Println(4 >> 2) // left shift bits by 2^i, i = 2 here

	// strings
	var str string = `hello`
	fmt.Println(len(str))

	fmt.Println(str[2])        // prints the byte value at 2 index
	fmt.Printf("%c\n", str[2]) // char value at 2 idnex

	fmt.Printf("%d\n", strings.Index("hello", "h"))
	fmt.Println(strings.ToLower("ABCD - anybody can dance"))
	fmt.Println(strings.Split("ABC D", " ")[0])
	fmt.Print(strconv.Itoa(12)) // int to string
	abc, _ := strconv.Atoi("hello")
	fmt.Println(abc) // string to int

	// time package
	t := time.Now()
	fmt.Printf("%d %d %d\n", t.Day(), t.Month(), t.Year())

	// pointers
	var ptr *int
	fmt.Println(ptr)

	var i1 = 5
	fmt.Println(&i1) // print the address of pointer

	// branching construct
	if !true {
		fmt.Println("yes true")
	} else {
		fmt.Println("yes false")
	}

	var n int = 2
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// calling func.
	res, err := sample_function(1, 2)
	fmt.Printf("%d %v\n", res, err.Error())

	// calling factorial

	fmt.Println(factorialIterative(4))
	fmt.Println(factorialRecursive(4))

	// switch
	var switchVar int = 1
	switch switchVar {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}
	// loop constructs , for loop
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("it is 5")
		}
	}
	// for range, used to go over collection
	// for i, val := range coll {} -> format
	for i, char := range "practice the golang" {
		fmt.Printf("position of char is %d & char is %c \n", i, char)
	}

	// assigning func to a var
	incr := inc
	fmt.Println(incr(1))

	// anonymous func. 
	anonymousFunc := func(x int, y int) int {return x * y}
	fmt.Println(anonymousFunc(2,3))
	
}


func inc(x int) int {return x} 

func sum(a int, b int) int { return a + b }

func sample_function(a int, b int) (int, error) {

	return 1, errors.New("error")
}

func factorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

func factorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)

}
