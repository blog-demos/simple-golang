package main

import (
    "net/http"
    "log"
    "fmt"
    "strings"
)

var times int

func main()  {
    var err error
    http.HandleFunc("/post/test-rute", parsePostParams)      //设置访问的路由
    err = http.ListenAndServe(":55551", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func parsePostParams(w http.ResponseWriter, r *http.Request)  {
    times++

    r.ParseForm() // 解析url传递的参数,对于POST则解析响应包的主体(request body)
    //注意:如果没有调用ParseForm方法,下面无法获取表单的数据
    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息

    a := strings.Join(r.Form["a"], "")
    b := strings.Join(r.Form["b"], "")
    c := strings.Join(r.Form["c"], "")

    fmt.Println("a", a)
    fmt.Println("b", b)
    fmt.Println("c", c)
    fmt.Println("times:", times)

    fmt.Fprintf(w, "调用次数：%d", times) //这个写入到w的是输出到客户端的
}
