package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
)

func main()  {
    http.HandleFunc("/post/test-rute", readContent)
    http.ListenAndServe(":55551", nil)
}

func readContent(w http.ResponseWriter, r *http.Request)  {
    defer r.Body.Close()

    con, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(con))

    fmt.Fprint(w, "Hello, Http.POST.")
}
