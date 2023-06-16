package ziface

type IServer interface {
	//啟動伺服器
	Start()
	//停止伺服器
	Stop()
	//運行伺服器
	Serve()
}
