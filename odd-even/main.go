package main

import "fmt"

func main() {
	num := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println(num)

	for _, numero := range num {
		if numero % 2 == 0 {
			fmt.Println("number", numero, "is even")
		}	else {
			fmt.Println("number", numero, "is odd")
		}
	}
}