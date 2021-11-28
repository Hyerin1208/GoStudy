package main

import "fmt"

func main() {
	k := 1

   // *의 개수 범위를 제한하는 변수
	q := 3

   // 공백의 개수 범위를 제한하는 변수
	for i := 0; i < 7; i++ {

		// 공백 부분
		for p := 0; p < q; p++ {
			fmt.Print(" ")
		}

		if i < 3 {
			q--
		} else {
			q++
		}
        
		// * 부분
		for j := 0; j < (2*k - 1); j++ {
			fmt.Print("*")
		}

		if i < 3 {
			k++
		} else {
			k--
		}

		fmt.Println()
	}
}