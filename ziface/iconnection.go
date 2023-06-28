package ziface

import "net"

type IConnection interface {
	//自動連線 讓當前的連線準備開始工作
	Start()

	//停止連線
	Stop()

	//獲取當前連線綁定的socket conn
	GetTCPConnection() *net.TCPConn

	//獲取當前連線模組的連線ID
	GetConnId() uint32

	//獲取遠端客戶的TCP狀態 ip port
	RemoteAddr() net.Addr

	//發送數據，將數據發送給遠端客戶
	Send(data []byte) error
}

// 定義一個處理連線業務的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
