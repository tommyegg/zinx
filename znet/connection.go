package znet

import (
	"fmt"
	"net"

	"github.com/tommyegg/zinx/utils"
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

	//該連線處理的方法router
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
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
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buff err", err)
			continue
		}

		//得到目前conn數據的request請求數據
		req := Request{
			conn: c,
			data: buf,
		}

		//從路由中，找到註冊綁定的Conn對應的router方法
		go func(request ziface.IRequest) {
			c.Router.PreHandler(request)
			c.Router.Handler(request)
			c.Router.PostHandler(request)
		}(&req)
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
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 獲取當前連線模組的連線ID
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

// 獲取遠端客戶的TCP狀態 ip port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 發送數據，將數據發送給遠端客戶
func (c *Connection) Send(data []byte) error {
	return nil
}
