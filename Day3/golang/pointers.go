package main

import "fmt"

func sayHello( msgPtr *string ) {

	//Dereferencing - the values stored at address pointed by msgPtrl will be printed here
	fmt.Println("Inside sayHello function ", *msgPtr )
	fmt.Println("Address pointed by msgPtr is ", msgPtr ) //Here the address pointed by pointer will be printed

	tmp := *msgPtr //The value stored at the address pointed by msgPtr is assigned to tmp string

	*msgPtr = tmp + " Golang" + " !"

	fmt.Println("Inside sayHello before return ", *msgPtr)

}

func main() {
   //msg is a string variable assigned with value "Hello"
   msg := "Hello"

   fmt.Println("Message before calling sayHello function is ", msg)
   fmt.Println("Address of msg string is ", &msg)

   //sayHello function is taking the address of msg string
   sayHello( &msg )

   fmt.Println("Message after calling sayHello function is ", msg )
}
