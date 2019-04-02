package main

import (
    "time"
    "fmt"
)

func main() {
    t1 := time.Date(2010, 9, 15, 0, 0, 0, 0, time.UTC);
    fmt.Println(t1)
    t2 := time.Now()
    fmt.Println(t2)

    fmt.Println(timeSub(t2, t1))
}

func timeSub(t1, t2 time.Time) int {
    t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
    t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

    return int(t1.Sub(t2).Hours() / 24)
}