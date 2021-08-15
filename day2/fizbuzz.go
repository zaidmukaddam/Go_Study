package main // for entry point of a go program

import "fmt" //importing the fmt library for Printf and Scanf

func main() {
	//for loop start
	for justhebuzz := 1; justhebuzz <= 100; justhebuzz++ {
		fizzthebuzz := "" // declaring a string variable "fizzthebuzz"

		if justhebuzz%3 == 0 {
			fizzthebuzz += "Fizz" // if justhebuzz is divisible by 3 append "Fizz" to the fizzthebuzz variable
		}

		if justhebuzz%5 == 0 {
			fizzthebuzz += "Buzz" // if justhebuzz is divisible by 5 append "Buzz" to the fizzthebuzz variable
		}

		if fizzthebuzz != "" {
			fmt.Printf("%s\n", fizzthebuzz) // if fizzthebuzz is not an empty string then print fizzthebuzz
		} else {
			fmt.Printf("%d\n", justhebuzz) // else print justhebuzz
		}
	}
	//for loop end
}
