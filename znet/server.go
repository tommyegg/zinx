package znet

import (
	"fmt"
	"net"

	"github.com/tommyegg/zinx/utils"
	"github.com/tommyegg/zinx/ziface"
)

// IServer的接口實現，定義一個server的伺服器模組
type Server struct {
	//伺服器的名稱
	Name string
	//伺服器綁定的ip版本
	IPVersion string
	//伺服器監聽的IP
	IP string
	//伺服器監聽的port
	Port int
	//當前的server增加一個router，server註冊的連線對應的業務處理
	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf(
		"[Zinx] Server Name: %s, listenner at IP: %s,Port:%d is starting\n",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s,MaxConn:%d,MaxPackeetSize:%d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start] Server Lintenner at IP: %s,Port %d", s.IP, s.Port)

	go func() {
		//1.取得一個tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("reslove tcp addr error:", err)
			return
		}

		//2.監聽伺服器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err:", err)
		}

		fmt.Println("start Zinx server succ,", s.Name, "succ, Lintenning...")
		var cid uint32
		cid = 0

		//3.阻塞等待客戶端連線，處理客戶端的業務（讀寫）
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err,", err)
				continue
			}

			dealConnection := NewConnection(conn, cid, s.Router)
			cid++

			go dealConnection.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	//啟動server的服務功能
	s.Start()

	//TODO 做一些啟動伺服器之後的額外流程

	//阻塞
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Successful")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}

	return s
}
