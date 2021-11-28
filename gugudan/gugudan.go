package main

import "fmt"

func main() {

	for dan := 1; dan <= 9; dan++ {

		// if dan == 3 {
		// 	continue
		// }

		fmt.Printf("%d단\n", dan)

		for j := 1; j <= 9; j++ {
// 3*2 만 출력되지 않도록
			if dan==3 && j == 2 {
				continue
			}

			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}

	}		
		
		fmt.Println()
	}