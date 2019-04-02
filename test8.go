package main

import "fmt"

var a = "G"

func main()  {
	n()
	m()
	n()
	y()
	n()
}

func n()  {
	fmt.Println(a)
}

func m()  {
	a := "O"
	fmt.Println(a)
}

func y()  {
	a = "Y"
	fmt.Println(a)
}