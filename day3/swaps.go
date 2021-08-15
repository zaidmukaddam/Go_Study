package main

import "fmt"

func swap1(a int, b int) (int, int) {
	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}

func swap2(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swap3(a int, b int) (int, int) {
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

func swap4(a int, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func swap5(a int, b int) (int, int) {
	return b, a
}

func main() {
	var first, second int

	fmt.Print("Enter first number : ")
	fmt.Scanln(&first)

	fmt.Print("Enter second number : ")
	fmt.Scanln(&second)

	first, second = swap1(first, second)
	first, second = swap2(first, second)
	first, second = swap3(first, second)
	first, second = swap4(first, second)
	first, second = swap5(first, second)

	fmt.Println("First number :", first)
	fmt.Println("Second number :", second)

}
