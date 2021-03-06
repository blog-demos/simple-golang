package main

import (
    "fmt"
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()                                                       // 解析url传递的参数,对于POST则解析响应包的主体(request body)
    // 注意: 如果没有调用ParseForm方法,下面无法获取表单的数据
    //fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
    //fmt.Println(r.Form["a"])
    //for k, v := range r.Form {
    //    fmt.Println("key:", k)
    //    fmt.Println("val:", strings.Join(v, ""))
    //}

    fmt.Fprintf(w, "Hello World.")
}

func main() {
    var err error
    http.HandleFunc("/test-rute", sayHello)      // 设置访问的路由
    err = http.ListenAndServe(":55550", nil) // 设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}