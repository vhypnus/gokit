package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("go debug.....")
    // 命令参数使用
    var argc = len(os.Args)
    var argv = append([]string{}, os.Args...)
    fmt.Printf("argc:%d\n", argc)
    fmt.Printf("argv:%v\n", argv)

    // 查看变量内存地址及值
    var aa = 1
    var bb = -1

    fmt.Println(aa)
    fmt.Println(bb)
}