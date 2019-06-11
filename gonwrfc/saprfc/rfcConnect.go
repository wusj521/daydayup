package saprfc

import (
	"fmt"
	//"reflect"
	//"testing"
	//"time"
	"gonwrfc/utils"

	"gonwrfc/gorfc"
	//"github.com/stretchr/testify/assert"
)

// var c *gorfc.Connection
var SAPconnection *gorfc.Connection

func abapSystem() gorfc.ConnectionParameter {
	// 读取配置文件STAR
	iniConfig := new(utils.Config)
	iniConfig.InitConfig("./config.ini")
	//直联机1参数
	iClient := iniConfig.Read("rfcConnect", "iClient") //client
	iUser := iniConfig.Read("rfcConnect", "iUser")     //SAP用户
	iPasswd := iniConfig.Read("rfcConnect", "iPasswd") //SAP密码
	iAshost := iniConfig.Read("rfcConnect", "iAshost") //SAP服务器IP

	return gorfc.ConnectionParameter{
		Dest:      "1test",
		Client:    iClient,
		User:      iUser,
		Passwd:    iPasswd,
		Lang:      "ZH",
		Ashost:    iAshost,
		Sysnr:     "00",
		Saprouter: "/H/123.125.21.51/H/",
		// Dest:      "1test",
		// Client:    "800",
		// User:      "wusj",
		// Passwd:    "xxxx",
		// Lang:      "ZH",
		// Ashost:    "192.168.100.66",
		// Sysnr:     "00",
		// Saprouter: "/H/123.125.21.51/H/",
		// Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
	}
}

func RfcConnect() error {
	//rfc con
	var err error
	SAPconnection, err = gorfc.ConnectionFromParams(abapSystem())
	if err == nil {
		fmt.Println("测试连接成功...1")
	} else {
		return err
	}
	// //call rfc
	// a := c.Alive()
	// if a == true {
	// 	return true
	// } else {
	// 	return false
	// }
	return nil
}

func RfcClose() error {
	//rfc con close
	err := SAPconnection.Close()
	if err != nil {
		return err
	}
	return nil
}
