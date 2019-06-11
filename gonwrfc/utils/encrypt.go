package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"time"
)

//md5方法
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Guid方法
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

//sha256
func Sha2(s string) string {
	h := sha256.New()

	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha512(s string) string {
	h := sha512.New()

	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 将2016-07-27 08:46:15这样的时间字符串转换时间戳
func StrToUnix(t string) int64 {
	// 先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
	// 后面就可以用time.Time对象的函数Unix进行获取
	//t2, err := time.Parse("2006-01-02 15:04:05", "2016-07-27 08:46:15")
	t2, err := time.Parse("2006-01-02", t)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(t2)
	//fmt.Println(t2.Unix())
	return t2.Unix()
	// output:
	//     2016-07-27 08:46:15 +0000 UTC
	//     1469609175
}
func UnixToStr(tunix int64) string {
	// 获取指定时间戳的年月日，小时分钟秒后 转为string
	t := time.Unix(tunix, 0)
	var date string

	date = fmt.Sprintf("%d-%d-%d\n", t.Year(), t.Month(), t.Day())
	return date
	//return date
	//t := time.Unix(1469579899, 0)
	//fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// output: 2016-7-27 8:38:19
}

func AddDay(tunix int64, nubmber int) string {
	// 对指定时间戳的年月日，进行增加和减少计算
	t := time.Unix(tunix, 0)
	d := t.AddDate(0, 0, -nubmber)
	days := d.Format("2006-01-02")
	//days = fmt.Sprintf("%d-%d-%d\n", d.Year(), d.Month(), d.Day())
	return days
}
func AddDay1(tunix int64, nubmber int) string {
	// 对指定时间戳的年月日，进行增加和减少计算
	t := time.Unix(tunix, 0)
	d := t.AddDate(0, 0, -nubmber)
	days := d.Format("20060102")
	//days = fmt.Sprintf("%d-%d-%d\n", d.Year(), d.Month(), d.Day())
	return days
}
func AddMonth(tunix int64, nubmber int) string {
	// 对指定时间戳的年月日，进行增加和减少计算
	t := time.Unix(tunix, 0)
	d := t.AddDate(0, -nubmber, 0)
	days := d.Format("20060102")
	//days = fmt.Sprintf("%d-%d-%d\n", d.Year(), d.Month(), d.Day())
	return days
}

func Getday(tunix int64, nubmber int) string {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(tunix, 0)
	var date string
	var day int
	if t.Day() < nubmber {
		day30 := t.Day() + 30
		mon := t.Month() - 1
		switch mon {
		case 3:
			mon = 03
		}
		day = day30 - nubmber
		date = fmt.Sprintf("%d-%d-%d\n", t.Year(), mon, day)
	} else if t.Day() == nubmber {
		mon := t.Month() - 1
		day := 28
		date = fmt.Sprintf("%d-%d-%d\n", t.Year(), mon, day)
	} else {
		day = t.Day() - nubmber
		date = fmt.Sprintf("%d-%d-%d\n", t.Year(), t.Month(), day)
	}

	//day := t.Day() - nubmber

	//date = fmt.Sprintf("%d-%d-%d\n", t.Year(), t.Month(), day)
	return date
	//return date
	//t := time.Unix(1469579899, 0)
	//fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// output: 2016-7-27 8:38:19
}
