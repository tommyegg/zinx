package znet

import (
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
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}

func NewServer(name string) ziface.IServer {

}
