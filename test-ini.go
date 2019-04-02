/*
@Time :     2018/10/25 15:22 
@Author :   Hickey Q
@File :     test-ini.go
@Software:  GoLand
*/

package main

import (
    "fmt"
    "strconv"
    "github.com/robfig/config"
)

type Student struct {
    name string
    age int
}

const (
    CFG_FIE_NAME = "students.ini"

    SECTION1 = "Student 1"
    SECTION2 = "Section 2"

    OPTION_NAME = "Name"
    OPTION_AGE = "Age"
)

func foo() {
    c := config.NewDefault()

    tom := Student{"Tom", 5}
    jerry := Student{"Jerry", 6}

    c.AddSection(SECTION1)
    c.AddOption(SECTION1, OPTION_NAME, tom.name)
    c.AddOption(SECTION1, OPTION_AGE, strconv.Itoa(tom.age))

    c.AddSection(SECTION2)
    c.AddOption(SECTION2, OPTION_NAME, jerry.name)
    c.AddOption(SECTION2, OPTION_AGE, strconv.Itoa(jerry.age))

    c.WriteFile(CFG_FIE_NAME, 0644, "All the students")

    fmt.Println("Done   ")
}

func bar() {
    c, err := config.ReadDefault(CFG_FIE_NAME)
    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    fmt.Println("Read student 1...")
    name, err := c.String(SECTION1, OPTION_NAME)
    if err != nil {
        fmt.Println("Get name failed:", err)
        return
    }

    age, err := c.Int(SECTION1, OPTION_AGE)
    if err != nil {
        fmt.Println("Get age failed:", err)
        return
    }

    fmt.Println("Student 1:", name, ",", age)

    fmt.Println("Read student 2...")
    name, err = c.String(SECTION2, OPTION_NAME)
    if err != nil {
        fmt.Println("Get name failed:", err)
        return
    }

    age, err = c.Int(SECTION2, OPTION_AGE)
    if err != nil {
        fmt.Println("Get age failed:", err)
        return
    }

    fmt.Println("Student 2:", name, ",", age)
}

func main() {
    //foo()
    bar()
}
