package main

import (
        "fmt"
        "log"
        "os"
        "plugin"
)


func main() {
        plug, err := plugin.Open("./printLog.so")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        log.Println("plugin opened")
        symLog, err := plug.Lookup("PrintLog")
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Printf("%T",symLog)   // 打印变量symLog的类型
        PrintLog, ok := symLog.(func())   // 类型断言,判断symLog是否为函数类型
        if !ok {
                fmt.Println("unexpected type from module symbol")
                os.Exit(1)
        }

        // 调用函数打印tomcat日志
        PrintLog()
}
