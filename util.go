package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
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

// parse 1442742699 to "2015-09-20 17:51:39"
func ToDatetime(ts int64) string {
	timeObj := time.Unix(ts, 0)
	timeStr := timeObj.Format(time.RFC3339)[0:19]

	temp := string(timeStr[0:10])
	temp += " "
	temp += string(timeStr[11:19])
	return temp
}

// parse "2012-09-01 00:00:00" to Time (+08:00)
func ParseTime(str string) (time.Time, error) {
	if len(str) != 19 {
		return time.Time{}, fmt.Errorf("time format error")
	}
	timestr := string([]byte(str)[0:10]) + "T" + string([]byte(str)[11:19]) + "+08:00"
	t, err := time.Parse(time.RFC3339, timestr)
	return t, err
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

func IntSlice(arr []int32) []int {

	var temp []int
	for _, v := range arr {
		temp = append(temp, int(v))
	}
	return temp
}

func IntSliceJoin(arr []int, sep string) string {
	var temp []string
	for _, v := range arr {
		temp = append(temp, fmt.Sprintf("%d", v))
	}
	return strings.Join(temp, sep)
}
func StrToMap(str, innerSep, outerSep string, urlDecode bool) (m map[string]string) {

	m = make(map[string]string)
	temp := strings.Split(str, outerSep)
	if len(temp) > 0 {

		for _, s := range temp {
			kv := strings.Split(s, innerSep)
			if len(kv) == 2 {
				k := kv[0]
				v := kv[1]
				if urlDecode {
					var err error
					k, err = url.QueryUnescape(k)
					if err != nil {
						continue
					}
					v, err = url.QueryUnescape(v)
					if err != nil {
						continue
					}

				}
				m[k] = v
			}
		}
	}
	return
}

func MapStrJoin(m map[string]string, innerSep string, outerSep string, urlEncode bool) string {
	var temp []string
	for k, v := range m {
		if urlEncode {
			k = url.QueryEscape(k)
			v = url.QueryEscape(v)
		}
		temp = append(temp, fmt.Sprintf("%s%s%s", k, innerSep, v))
	}
	return strings.Join(temp, outerSep)
}

func UrlJoin(m map[string]string) string {
	return MapStrJoin(m, "=", "&", true)
}

func UrlSplit(str string) (m map[string]string) {

	return StrToMap(str, "=", "&", true)

}

func Btoi(b bool) int {

	if b {
		return 1
	}
	return 0
}

func JsonOut(inter interface{}) []byte {
	res, err := json.Marshal(inter)
	if err != nil {
		return []byte("")

	}
	return res
}

func Printf(format string, args ...interface{}) (n int, err error) {
	origin_msg := fmt.Sprintf(format, args...)
	return fmt.Printf("[%s]%s\n", CurTime(), origin_msg)
}

func StringToInterface(old []string) []interface{} {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = interface{}(v)
	}
	return new
}

var (
	mapStrIntfTyp = reflect.TypeOf(map[string]interface{}(nil))
)

func Msgpack_Unpack(buffer []byte, v interface{}) (err error) {

	var mh codec.MsgpackHandle
	mh.MapType = mapStrIntfTyp

	dec := codec.NewDecoderBytes(buffer, &mh)
	err = dec.Decode(&v)
	return err
}

func Msgpack_Pack(buffer *[]byte, v interface{}) (err error) {

	var mh codec.MsgpackHandle
	mh.MapType = mapStrIntfTyp

	enc := codec.NewEncoderBytes(buffer, &mh)
	err = enc.Encode(v)

	return err
}

func main() {
	fmt.Println(CurDate())
	fmt.Println(CurTime())
}
