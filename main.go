package main

import "fmt"

func main() {
	f(9)

}

func f(a int) {

	if a <= 2 {
		fmt.Println(a)
	} else {
		f(a / 2)
		fmt.Println(a)
	}

}
