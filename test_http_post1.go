package main

import (
    "net/http"
    "log"
    "fmt"
    "strings"
)

func main()  {
    var err error
    http.HandleFunc("/post/test-rute", parseFormParams)
    err = http.ListenAndServe(":55551", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func parseFormParams(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    a := strings.Join(r.Form["a"], "")
    fmt.Fprintf(w, "参数 a = %s", a) //这个写入到w的是输出到客户端的
}
