package ziface

/*
	IRequest接口
	把客戶端請求訊息的鏈結，和請求的數據包裝到一個Request中
*/

type IRequest interface {
	//得到當前鏈結
	GetConnection() IConnection
	//得到請求的訊息
	GetData() []byte
	//得到請求訊息的ID
	GetMsgId() uint32
}
