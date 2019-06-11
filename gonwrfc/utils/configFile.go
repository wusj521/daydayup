//配置文件工具类
//
package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const middle = "====="

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	value, found := c.Mymap[key]
	if !found {
		return ""
	}
	return value
}

func (c Config) ReadInt(node, key string) int {
	key = node + middle + key
	value, found := c.Mymap[key]
	if !found {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}

// 读取配置文件
// iniConfig := new(util.Config)
// iniConfig.InitConfig("./ftp_upload.ini")
// host := iniConfig.Read("ftp", "host")             // ftp地址
// port := iniConfig.Read("ftp", "port")             // 端口号
// username := iniConfig.Read("ftp", "username")     // 用户名
// password := iniConfig.Read("ftp", "password")     // 密码
// local_dir := iniConfig.Read("ftp", "local_dir")   // 本地目录
// remote_dir := iniConfig.Read("ftp", "remote_dir") // 远程目录

// var ftpserverAddr = host + ":" + port
