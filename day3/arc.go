package main

import "fmt"

func main() {
	var rad float64         // a float64 variable
	const PI float64 = 3.14 // Constant
	var area float64        // a float64 variable to store area
	var ci float64          // a float64 variable to store circumfernce

	fmt.Print("Enter radius of Circle : ")
	fmt.Scanln(&rad) 
	//Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.

	area = PI * rad * rad // calutate area

	fmt.Println("area : ", area) 
	// Println formats using the default formats for its operands and writes to standard output.

	ci = 2 * PI * rad //calculate circumference
	fmt.Println("circumference : ", ci)
}
