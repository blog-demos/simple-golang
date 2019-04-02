package main

import ("fmt"; "./Package2")

func main()  {
	var sum int = Package2.Add(3, 4)
	fmt.Println(sum)
}

