package main

import (
    "fmt"
    "time"
)

// 当前日期
func CurDate() string {

    return time.Now().Format(time.RFC3339)[0:10]
}

// 当前日期时间
func CurTime() string {

    timeStr := time.Now().Format(time.RFC3339)[0:19]

    temp := string(timeStr[0:10])
    temp += " "
    temp += string(timeStr[11:19])
    return temp
}


func main() {
    fmt.Println(CurDate())
    fmt.Println(CurTime())
}
