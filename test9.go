package main

import "fmt"

var a1 string

func main()  {
	a1 = "G"
	fmt.Println(a1)
	m1()
}

func m1() {
	a1 := "O"
	fmt.Println(a1)
	m2()
}

func m2()  {
	fmt.Println(a1)
}