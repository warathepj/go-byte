package main

import "fmt"

func main() {
	// Declaring a byte variable
	var myByte byte = 65 // 65 is the ASCII value for 'A'

	// Printing the byte's integer value
	fmt.Println("The integer value of myByte is:", myByte)

	// Printing the byte as a character
	fmt.Printf("The character value of myByte is: %c\n", myByte)

	// Declaring a byte using a character literal
	anotherByte := 'B' // Character literals are bytes (if they are ASCII) or runes

	// Printing the second byte's integer value
	fmt.Println("The integer value of anotherByte is:", anotherByte)

	// Printing the second byte as a character
	fmt.Printf("The character value of anotherByte is: %c\n", anotherByte)

	// A byte slice (a sequence of bytes)
	byteSlice := []byte{72, 101, 108, 108, 111} // ASCII values for "Hello"

	// Printing the byte slice as a slice of integers
	fmt.Println("The byte slice as integers:", byteSlice)

	// Converting a byte slice to a string
	fmt.Println("The byte slice as a string:", string(byteSlice))

	// Converting a string to a byte slice
	myString := "World"
	stringToBytes := []byte(myString)
	fmt.Println("The string 'World' as a byte slice:", stringToBytes)
}
