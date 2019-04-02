package main

import (
    "fmt"
    "github.com/golang/protobuf/proto"
    tp "./testproto"
)

func main() {
    t := &tp.Tel{
        Number:     2147483647,
    }

    p := &tp.Person{
        Name:   "Bob abc",
        Age:    18,
        Tel:    t,
    }

    s := &tp.Student{
        Person: p,
        SchoolId:   10001,
    }

    out_t, _ := proto.Marshal(t)
    out_p, _ := proto.Marshal(p)
    out_s, _ := proto.Marshal(s)

    fmt.Println(out_t)
    fmt.Println(out_p)
    fmt.Println(out_s)
    fmt.Println(string(out_s))
}
