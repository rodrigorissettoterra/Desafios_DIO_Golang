package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println(i, "PIN-PAN")
		}
		if i%3 == 0 {
			fmt.Println(i, "pin")
		}
		if i%5 == 0 {
			fmt.Println(i, "pan")
		}
		if (i%3 != 0) && (i%5 != 0) {
			fmt.Println(i)
		}
	}
}
