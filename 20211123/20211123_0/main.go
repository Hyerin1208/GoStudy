package main

import (
	"fmt"
	"math"
)

const epsilon = 0.000001 // 매우작음

func equal(a, b float64) bool {

	return math.Nextafter(a, b) == b
	// if a > b {
	// 	if a-b <= epsilon {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }else {
	// 	if b-a <= epsilon {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
}

func main() {

	// var a int = 10
	// var b int = 20
	// var f float64 = 123231255.546

	// fmt.Println(1)

	// fmt.Printf("a : %d\\ b: %d\t f: %f\n", a, b, f)

	/*
	%T : 데이터 타입 출력
	%t : 불리언을 true/ false
	%d : 10진수 정수값으로 출력
	%b : 2진수 출력
	%c : 유니코드 문자를 출력
	%x : 16진수로 출력 (소)
	%X : 16진수로 출력 (대)
	%e : 지수 형태로 실수값 출력
	%s : 문자열을 출력
	%p : 메모리 주소값을 출력
	*/

	var str string = "H\nllo\tWor\\l\"d\""
	fmt.Println(str)

	var input int32
	var input1 int32
	// 입력을 받는다. 인자로는 해당 변수의 메모리주소.
	fmt.Scan(&input, &input1)

	// 비트연산 AND, OR, XOR, 비트 클리어
	// &, |, ^, &^(특정비트를 0으로)
	// 시프트 연산 <<(왼쪽), >>(오른쪽)

	// 00000010 << 1

	/*
	AND 1일 때만 1
	A 	B	A&B
	-------------
	0	0	0
	1	0	0
	0	1	0
	1	1	1
	*/
	// 10 & 34 = 2
	/*
	0000 1010 10
	0010 0010 34
	0000 0010 2
	*/

	/*
	OR : 하나라도 1이면 1
	0	0	0
	1	0	1
	0	1	1
	1	1	1

	0000 1010		10
	0010 0010		34
	0010 1010		42
	*/

	/*
	A^B
	XOR(A와 B가 다르면 1)
		0	0	0
		1	0	1
		0	1	1
		1	1	0

	0000 1010		10
	0010 0010		34
	0010 1000		40
	*/

	/*
	== , !=, <, >, <=, >=
	*/

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// fmt.Printf("%f+%f==%f : %v/n", fa, fb, fc, fa+fb == fc)
	// fmt.Println(fa + fb)

	fmt.Println("%0.18f + %0.18f = %0.18f \n", a, b, a+b)
	fmt.Printf("%0.18f ==%0.18f : %v\n", c, a+b, equal(a+b, c))

	// &&, ||, !, ++, --

	// []
	// . 
	// & 
	// * 
	// ...
	// :

	var test int8 = 30
	test <<= 2
	a += 8
	println(test)

}