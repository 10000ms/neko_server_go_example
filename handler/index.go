package handler

import (
    "fmt"
    "log"
    "strings"

    "neko_server_go"
)

func Index(c *neko_server_go.Context, w neko_server_go.ResWriter) {
    err := c.Request.ParseForm() //解析参数，默认是不会解析的
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Println(c.Request.Form) //这些信息是输出到服务器端的打印信息
    fmt.Println("path", c.Request.URL.Path)
    fmt.Println("scheme", c.Request.URL.Scheme)
    fmt.Println(c.Request.Form["url_long"])
    for k, v := range c.Request.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    _, err = fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
    if err != nil {
        log.Fatal(err)
        return
    }
}
