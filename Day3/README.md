# Day 3

## Lab - User-defined golang functions
Create a file named functions.go with the below content
```
package main

import "fmt"

func yetAnotherFunction() {
	fmt.Println("Yet Another Function invoked")
}

func main() {
   fmt.Println( sayHello("Golang") ) 
   fmt.Println( sayHello("World") ) 
   yetAnotherFunction()
}

// This function accepts a string input and returns a string output
func sayHello( msg string ) string {
  return "Hello, " + msg + " !"
}


/* function overloading is not supported by golang
func sayHello( ) string {
  return "Hello World !"
}
*/
```

Run it
```
go run ./functions.go
```

Expected output
![image](https://github.com/user-attachments/assets/9901da2f-046b-4848-91e5-d5320af8ab6c)

## Lab - Function with multiple returns
function-with-multiple-returns.go
```
package main

import "fmt"

func myFunction() (int,int) {
  return 10, 20
}

func main() {
   x, y := myFunction() // := is a short form of declaring a new variable and initializing some value 

   fmt.Println ( "Value of x is ", x )
   fmt.Println ( "Value of y is ", y )
}
```

Run it
```
go run ./functions-with-multiple-returns.go
```

Expected output
![image](https://github.com/user-attachments/assets/95e060bc-55cb-4ff8-8f44-3fa87fcc6fb2)
