/*
@Time :     2018/9/4 15:17 
@Author :   Hickey Q
@File :     test_template.go
@Software:  GoLand
*/

package main

import (
    "html/template"
    "os"
)

type Person struct {
    UserName string
}

func main(){
    t := template.New("Template Sample.")
    t, _ = t.Parse("user_name {{.UserName}}")
    p := Person{ UserName:"Alice" }
    t.Execute(os.Stdout, p)
}
