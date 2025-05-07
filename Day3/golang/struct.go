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

   var rectPtr *Rectangle
   rectPtr = &rectangle

   fmt.Println("Length : ", rectPtr.length)
   rectPtr.length = 500
   fmt.Println("Length : ", rectPtr.length)

   fmt.Println("Rectangle Length : ", rectangle.length )
}
