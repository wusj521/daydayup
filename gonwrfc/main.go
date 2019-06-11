package main

import (
	"gonwrfc/saprfc"
	"log"
)

func main() {
	var err error
	err = saprfc.RfcConnect()
	if err != nil {
		log.Println("sapconnect err:", err)
	} else {
		log.Println("sapconect  is ok.")
	}

	err = saprfc.Bank2sap()
	if err != nil {
		log.Println("sap call err:", err)
	} else {
		log.Println("sap call is ok.")
	}

	//Close connect SAP
	err = saprfc.RfcClose()
	if err != nil {
		log.Println("关闭连接失败：", err)
	} else {
		log.Println("关闭连接-成功！", err)
	}
}
