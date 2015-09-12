package main

import (
    "fmt"
    "time"
    "crypto/md5"
    "strconv"
    "io"
    "encoding/hex"
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

// 一致性哈希
func ConsistentHash(key string, count int) int {
    hash := Substr(Md5("-"+key), 8, 5)
    value, _ := strconv.ParseInt(hash, 16, 32)
    return int(value) % count
}

// md5摘要，返回16进制编码过的数据
func Md5(str string) string {
    h := md5.New()
    io.WriteString(h, str)
    return hex.EncodeToString(h.Sum(nil))
}

// md5摘要，返回raw data
func Md5Raw(str string) string {
    h := md5.New()
    io.WriteString(h, str)
    return string(h.Sum(nil)) // 返回rawdata的字符串格式
}

// 截取子串
func Substr(str string, start, count int) string {
    ByteArr := []byte(str) // 字符串转字节数组
    len := len(ByteArr)
    min := start + count
    if min > len {
        min = len
    }
    return string(ByteArr[start:min])
}



func main() {
    fmt.Println(CurDate())
    fmt.Println(CurTime())
}
