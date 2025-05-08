package main

import  ( 
          "os"
)

func createFile( name string, content string ) {

	myFile, err := os.Create(name)

	if err != nil {
          panic(err)
	}

	myFile.WriteString(content) 
	myFile.Sync() //This will ensure the file content is flushed out to disk

}

func main() {
	createFile( "myfile.txt", "This is a test file created using golang" )	
}
