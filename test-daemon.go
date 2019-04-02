/*
@Time :     2018/10/25 16:23 
@Author :   Hickey Q
@File :     test-daemon.go
@Software:  GoLand
*/

package main

import (
    "fmt"
    "time"
)

func main()  {
    for {
        var h, m, s = time.Now().Clock()

        fmt.Printf("%d:%d:%d\n", h, m, s)
        time.Sleep(1 * time.Second)
    }
}
