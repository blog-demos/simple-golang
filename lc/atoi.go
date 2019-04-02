/*
@Time :     2018/9/14 16:47 
@Author :   Hickey Q
@File :     atoi.go
@Software:  GoLand
*/

package main

import "fmt"

func myAtoi(str string) int {


    for i, elem := range str{
        if ' ' == elem {
            continue;
        }

        fmt.Println(i, str[i], string(elem))
    }

    return 0;
}

func main()  {
    myAtoi(" -987");
}
