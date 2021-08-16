package main

import "fmt"

func lcm(t1 int, t2 int) {
	var l int = 1
	if t1 > t2 {
		l = t1
	} else {
		l = t2
	}
	for {
		if l%t1 == 0 && l%t2 == 0 {

			fmt.Printf("LCM of %d and %d is %d", t1, t2, l)
			break
		}
		l++
	}
}

func gcd(t1 int, t2 int) {
	var gcdnum int

	for i := 1; i <= t1 && i <= t2; i++ {
		if t1%i == 0 && t2%i == 0 {
			gcdnum = i
		}
	}
	fmt.Printf("GCD of %d and %d is %d", t1, t2, gcdnum)

}

func main() {
	var n1, n2, action int
	fmt.Println("Enter two positive integers : ")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	fmt.Println("Enter 1 for LCM and 2 for GCD")
	fmt.Scanln(&action)

	switch action {
	case 1:
		lcm(n1, n2)
	case 2:
		gcd(n1, n2)
	}
}
