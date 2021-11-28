package main

import "fmt"

func main() {

	for dan := 1; dan <= 9; dan++ {

		fmt.Printf("%d단\n", dan)

		for j := 1; j <= 9; j++ {
			if j == 5 {
				continue
			}
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}
		}
		
		fmt.Println()
	}