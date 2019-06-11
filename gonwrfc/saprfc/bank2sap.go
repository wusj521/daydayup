//CALL SAP RFC 以结构方式传值，可以实现通过RFC写入SAP DB表中或从SAP中读取数据。
//2018/06/03 By wusj

package saprfc

import (
	"fmt"
	//"reflect"
	//"testing"
	//"time"
	//"github.com/sap/gorfc/gorfc"
	//"github.com/stretchr/testify/assert"
)

////连接参数
// func abapSystem() gorfc.ConnectionParameter {
// 	return gorfc.ConnectionParameter{
// 		Dest:      "1test",
// 		Client:    "800",
// 		User:      "wusj",
// 		Passwd:    "wusjcxy",
// 		Lang:      "ZH",
// 		Ashost:    "192.168.100.66",
// 		Sysnr:     "00",
// 		Saprouter: "/H/123.125.21.51/H/",
// 		// Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
// 	}
// }

func Bank2sap(r []string) error {

	// c, err := gorfc.ConnectionFromParams(abapSystem())
	// if err == nil {
	// 	fmt.Println("测试连接...1")
	// }

	//c.Close()
	//q := c.Alive()
	//fmt.Println(q)
	//o := c.Open()
	// fmt.Println(o)
	// if o != nil {

	// 	c, err := gorfc.ConnectionFromParams(abapSystem())
	// 	if err == nil {
	// 		fmt.Println(c)
	// 		fmt.Println("测试连接...2")
	// 	}
	// }
	//r, err := c.Call("ZWZK_TM_99", params)

	//单参数传值方式
	// params := map[string]interface{}{
	// 	"NUMBER":  banfn,
	// 	"REL_CODE":    prgpr,
	// }

	////判断数据类型合规性
	// fmt.Println("判断数据类型合规性", r)
	if r[10] == "" {
		r[10] = "0" //交易金额
	}
	if r[11] == "" {
		r[11] = "0" //余额
	}
	if r[12] == "" {
		r[12] = "0" //可用余额
	}
	// fmt.Println("判断数据类型合规性", r)
	////判断数据类型合规性

	//结构体方式传参数
	//params := map[string]interface{}{"I_ZRFC_DATA01": map[string]interface{}{}}
	//SAP RFC结构体方式传参数
	params := map[string]interface{}{ //结构体方式传参数
		"R_BANKSAP": map[string]interface{}{ //向RFC中的tables试图下传值失败，目前传值在Import 是Okay

			"STATUS":     r[0],  //状态‘0’:正常‘1’: 已经被抹账
			"JYDATE":     r[1],  //交易日期
			"JYTIME":     r[2],  //交易时间
			"YWULEIX":    r[3],  //业务类型0系统内手工转1系统内自动转2对外支付3其他
			"LIUHAO":     r[4],  //流水号
			"LIUSXHAO":   r[5],  //流水序号
			"ZHANGHAO":   r[6],  //账号
			"HUMING":     r[7],  //户名
			"SHOUZBS":    r[8],  //收支标志‘C’:贷‘D’:借
			"BIZHONG":    r[9],  //币种
			"JYEDU":      r[10], //交易金额
			"YUE":        r[11], //余额
			"KEYEDU":     r[12], //可用余额
			"DFZHANGHAO": r[13], //对方账号
			"DFHUMING":   r[14], //对方户名
			"DFDIZ":      r[15], //对方地址
			"DFKHHAO":    r[16], //对方开户行行号
			"DFKHMING":   r[17], //对方开户行行名
			"PIAOJZL":    r[18], //票据种类
			"PIAOJHM":    r[19], //票据号码
			"PIAOJMING":  r[20], //票据名称
			"PIAOJDATE":  r[21], //票据签发日期
			"FUYAN":      r[22], //附言
			"BEIZHU":     r[23], //备注

		},
	} //params 赋值

	// //rfc con
	// c, err := gorfc.ConnectionFromParams(abapSystem())
	// if err == nil {
	// 	fmt.Println("测试连接...1")
	// } else {
	// 	return err
	// }
	// fmt.Println(params)

	//call rfc.//rfc connection succeed if alive is ture else connection bad.
	//c := SAPconnection.Alive()
	//fmt.Println(SAPconnection.Alive())
	if SAPconnection.Alive() == true {
		fmt.Println("SAP已经连接成功！")
		_, err := SAPconnection.Call("Z_WUSJ_TEST02", params)
		if err != nil {
			return err
			fmt.Println("CALL RFC错误提示：", r, err)
		}
	}
	if SAPconnection.Alive() != true {
		fmt.Println("SAP重新连接中...")
		err := RfcConnect()
		if err != nil {
			fmt.Println("错误提示：", r, err)
			return err
		} else {
			fmt.Println("SAP重新连接成功！")
		}
	}
	//返回数据处理,接收tables视图下参数问题，其它export changing没有测试过

	// echoStruct := r["RETURN_FLAG"].([]interface{})

	// for _, value := range echoStruct {
	// 	values := value.(map[string]interface{})
	// 	// 	//		fmt.Println(len(values)) //打印行数
	// 	// 	//		fmt.Println(values["MATNR"])//打印某个字段的值
	// 	// 	//	delete(values, "MAKTX")
	// 	// 	//		fmt.Println(values["MATNR"])
	// 	// 	//		fmt.Println(values["MATNR"])

	// 	// 	//ebeln := values["EBELN"]
	// 	fmt.Println(value["RETURN_FLAG"])
	// 	// fmt.Println(value["MATERIALCODE"])

	// } //循环RFC返回值

	//c.Close()

	//c.Close()
	return nil
}
