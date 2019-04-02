/*
@Time :     2018/9/7 10:34 
@Author :   Hickey Q
@File :     zig-zag.go
@Software:  GoLand
*/

package main

import "fmt"

func convert(s string, numRows int) string {
    for i, elem := range s{
        fmt.Println(i, s[i], string(elem))
    }

    return ""
}

func main()  {
    s := "PAYPALISHIRING"
    numRows := 3

    result := convert(s, numRows)

    fmt.Println(result)
}

