package znet

import (
	"errors"
	"fmt"
	"io"
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
		// buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("recv buff err", err)
		// 	continue
		// }

		//創建一個拆包解包的對象
		dp := NewDataPack()
		//讀取客戶端的Msg Head 二進制流 8個字節
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error")
			break
		}

		//拆包，得到MsgId 和 msgDatalen 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		//根據datalen 再次讀取Data, 放在 msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetData(data)

		//得到目前conn數據的request請求數據
		req := Request{
			conn: c,
			msg:  msg,
		}

		//從路由中，找到註冊綁定的Conn對應的router方法
		go func(request ziface.IRequest) {
			c.Router.PreHandler(request)
			c.Router.Handler(request)
			c.Router.PostHandler(request)
		}(&req)
	}
}

// 提供一個SendMsg方法，將要發送給客戶端的數據先進行封包
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	//將data進行封包 MsgDataLen|MsgId|Data
	dp := NewDataPack()
	//MsgDataLen|MsgId|Data
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id=", msgId)
		return errors.New("Pack error msg")
	}

	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write msg id", msgId, "error:", err)
		return errors.New("conn write err")
	}

	return nil
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
