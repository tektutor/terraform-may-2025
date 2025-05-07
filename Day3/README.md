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

## Lab - Golang loops

loops.go
```
ackage main

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
```

Run it
```
go run ./loops.go
```
Expected output
![image](https://github.com/user-attachments/assets/9144de70-58d4-498a-9b0a-2e2b6248daf9)


## Lab - Golang map

maps.go
```
package main

import "fmt"

func main() {

	toolsPath := map[string]string {
           "java_home": "/usr/lib/jdk11",
	   "mvn_home" : "/usr/share/maven",
	}

	fmt.Println("Java Home Directory :", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
		fmt.Println(key,value)
	}

	//delete a key-value pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)
}
```

Run it
```
go run ./maps.go
```

Expected output
![image](https://github.com/user-attachments/assets/312b3bbf-cd08-4305-9b75-ab253f877b51)
![image](https://github.com/user-attachments/assets/3694d95c-ac28-4e91-8f76-e20fa62d866a)


## Lab - Golang pointers
pointers.go
```
package main

import "fmt"

func sayHello( msgPtr *string ) {

	//Dereferencing - the values stored at address pointed by msgPtrl will be printed here
	fmt.Println("Inside sayHello function ", *msgPtr )
	fmt.Println("Address pointed by msgPtr is ", msgPtr ) //Here the address pointed by pointer will be printed
	fmt.Println("Address pointed by msgPtr is ", &msgPtr ) //Prints the address of msgPtr

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
```

Run it
```
go run ./pointers.go
```

Expected output
![image](https://github.com/user-attachments/assets/bb48441e-325b-46f0-8dd1-da7ea8e99e27)
![image](https://github.com/user-attachments/assets/0db5e7d4-8544-4ed7-9a5a-94aa9e087ecf)

## Lab - golang struct with methods
struct.go
```
package main

import "fmt"

type Rectangle struct {
   length int
   width  int
}

//golang Method of struct Rectangle
func (rect Rectangle) Area() int {
    area := rect.length * rect.width
    return area
}

func (rect Rectangle) GetLength() int {
	return rect.length
}

func (rect Rectangle) GetWidth() int {
	return rect.width
}

func main() {
   rectangle := Rectangle {
      length: 100,
      width : 200, 
   }

   fmt.Printf("Length of rectangle : %d\n", rectangle.GetLength())
   fmt.Printf("Width of rectangle  : %d\n", rectangle.GetWidth())
   fmt.Printf("Area of rectangle   : %d\n", rectangle.Area())
}
```

Run it
```
go run ./struct.go
```

Expected output
![image](https://github.com/user-attachments/assets/eda23e44-cae8-4018-afb0-f56af0d75c9b)

## Lab - Golang switch case
```
package main

import "fmt"

func main() {

	var direction string

	fmt.Print("Possible values are east,west,south,north")

	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {
	case "east":
		fmt.Println("You entered direction ", direction)
	case "west":
		fmt.Println("You entered direction ", direction)
	case "south":
		fmt.Println("You entered direction ", direction)
	case "north":
		fmt.Println("You entered direction ", direction)
	default:
		fmt.Println("Invalid direction", "possible values are east, west, north, south")
	}
}
```

Run it
```
go run ./switch-case.go
```

Expected output
![image](https://github.com/user-attachments/assets/a6aaa1a7-291e-47e7-90e8-674de63b092e)
