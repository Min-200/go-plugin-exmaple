package main

import (
        "bytes"
        "fmt"
        "os/exec"
)

func PrintLog() {
        // 待执行的Shell命令
        cmdStr := "cat /home/tomcat/logs/catalina.out"
        // 创建执行命令的对象
        cmd := exec.Command("sh", "-c", cmdStr)
        // 执行命令并捕获输出
        output := bytes.NewBuffer(nil)
        cmd.Stdout = output
        cmd.Stderr = output
        // 启动命令并捕获错误
        err := cmd.Start()
        if err != nil {
                fmt.Printf("执行命令错误")
        }
        // 等待命令执行完成
        err = cmd.Wait()
        if err != nil {
                fmt.Printf("执行报错")
        }else {
                fmt.Printf(output.String())
        }
}
