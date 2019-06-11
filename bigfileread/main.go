// bigfileread project main.go
// by wusj 20190610
//读取大文件GB级别的，逐行读取, 一行是一个[]byte, 多行就是[][]byte
package main

import (
	"bufio"
	"fmt"

	//	"io/ioutil"
	"os"
	"time"
)

// 逐行读取, 一行是一个[]byte, 多行就是[][]byte
func readByLine(filename string) (lines [][]byte, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)
	for {
		line, _, err := bufReader.ReadLine() //按行读
		if err != nil {
			err = nil
			break
		} else {
			lines = append(lines, line)
		}
	}
	return
}

func main() {
	// fmt.Println("Hello World!")
	//计算运行开始时间
	t1 := time.Now() // get current time

	filename := "D:\\test\\testbackcsv\\listdata1.csv"
	// filename := "D:\\test\\testbackcsv\\lims.txt"
	lines, err := readByLine(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	//计算运行结束时间
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	// no := 0
	// for k, _ := range lines { //多行byte数据
	// 	str := string(lines[k]) //每行数据从byte 转为string
	// 	fmt.Println(str)
	// 	no = no + 1
	// 	fmt.Println("Hangshu are:", no)
	// }
	// // str := string(lines[0])
	// // fmt.Println(str)

}
