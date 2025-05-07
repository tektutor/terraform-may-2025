package main

import "fmt"

func main() {

	count := 5 //Declares a count variant of type int and assigns a value 5

	for count > 0 {
	   fmt.Println("Before decrementing ", count)

	   count-- //equivalent to count = count - 1

	   fmt.Println("After decrementing ", count)

	   //--count pre-decrement is not supported in golang unlike C++
	   //++count pre-increment is not supported in golang unlike C++
	   //count++ is supported in golang
	}
	fmt.Println("Value of count is", count, " after for loop")

	count = 0 //Variable is already declared, in this line we are just resetting the count value to 0

	for count=1; count<10; count++ {
	   fmt.Println( count )
	}
}
