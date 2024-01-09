package main

import "fmt"

func main() {
	var size int

	fmt.Scanf("%d", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
				fmt.Printf("#")
			} else {
				fmt.Print(" ")
			}
			if j == size-1 {
				fmt.Printf("\n")
			}
		}
	}
}
