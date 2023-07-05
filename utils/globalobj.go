package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tommyegg/zinx/ziface"
)

/*
	存放框架全域變數
*/

type GlobalObj struct {
	//server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	//zinx
	Version        string
	MaxConn        int    //允許最大連線數
	MaxPackageSize uint32 //允許數據包的最大值
}

/*
	定義一個全域的對外Globalobj
*/

var GlobalObject *GlobalObj

// 從zinx.json去讀取用戶自訂的參數
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//將json資料解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 提供一個init方法，初始化當前的GlobalObject
func init() {
	//如果客戶端沒有設定檔的預設值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
