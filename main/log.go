package main

import (
    "log"
    "os"
)

// 写日志案例

func main() {
    // 定义一个文件
    fileName := "ll.log"
    logFile, err := os.Create(fileName)
    defer logFile.Close()
    if err != nil {
        log.Fatalln("open file error !")
    }

    // 创建一个日志对象
    debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
    debugLog.Println("A debug message here")

    //重新配置一个日志格式的前缀
    debugLog.SetPrefix("[Info]")
    debugLog.Println("A Info Message here ")

    //重新配置log的Flag参数
    debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
    debugLog.Println("A different prefix")
}
