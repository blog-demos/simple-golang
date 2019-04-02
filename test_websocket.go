/*
@Time :     2018/9/5 16:15 
@Author :   Hickey Q
@File :     test_websocket.go
@Software:  GoLand
*/

package main
import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
)

func Echo(ws *websocket.Conn) {
    var err error
    for {
        var reply string
        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive", err)
            break
        }

        fmt.Println("Received back from client: " + reply)
        msg := "Received: " + reply
        fmt.Println("Sending to client: " + msg)
        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
}

func main() {
    http.Handle("/", websocket.Handler(Echo))
    if err := http.ListenAndServe(":55551", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}