package ziface

type IServer interface {
	//啟動伺服器
	Start()
	//停止伺服器
	Stop()
	//運行伺服器
	Serve()
	//路由功能，給當前的伺服器註冊一個路由方法，給客戶端使用
	AddRouter(router IRouter)
}
