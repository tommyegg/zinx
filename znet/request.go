package znet

import "github.com/tommyegg/zinx/ziface"

type Request struct {
	//已經和客戶端建立好的鏈結
	conn ziface.IConnection
	//客戶端請求的數據
	msg ziface.Imessage
}

// 得到當前鏈結
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 得到請求的訊息
func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
