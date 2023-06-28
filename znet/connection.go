package znet

import (
	"fmt"
	"net"

	"github.com/tommyegg/zinx/ziface"
)

type Connection struct {
	//當前練線的socket tcp 套接字
	Conn *net.TCPConn

	//當前連線的id
	ConnId uint32

	//當前的連線狀態
	isClosed bool

	//當前練線所綁定的處理業務方法API
	handleApi ziface.HandleFunc

	//告知當前連線已退出的channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connID,
		handleApi: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}

	return c
}

// 連線的讀業務方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ...")
	defer fmt.Println("connID=", c.ConnId, "Reader is exit, remote addr is:", c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		//讀取客戶端的數據到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buff err", err)
			continue
		}

		//調用目前連線所綁定的HandelAPI
		if err := c.handleApi(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnId, "handle is error", err)
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnId=", c.ConnId)
	//啟動從當前連線的讀業務

	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn stop().. ConnID=", c.ConnId)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//關閉socket 連線
	c.Conn.Close()

	//回收資源
	close(c.ExitChan)
}

// 獲取當前連線綁定的socket conn
func GetTCPConnection(c *Connection) *net.TCPConn {
	return c.Conn
}

// 獲取當前連線模組的連線ID
func GetConnId(c *Connection) uint32 {
	return c.ConnId
}

// 獲取遠端客戶的TCP狀態 ip port
func RemoteAddr(c *Connection) net.Addr {
	return c.Conn.RemoteAddr()
}

// 發送數據，將數據發送給遠端客戶
func Send(data []byte) error {
	return nil
}
