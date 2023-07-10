package ziface

/*
	將請求的訊息封裝到一個Message中，定義抽象的接口
*/

type Imessage interface {
	//取得訊息的id
	GetMsgId() uint32
	//取得訊息的長度
	GetMsgLen() uint32
	//取得訊息的內容
	GetMsgData() []byte

	//設定訊息的id
	SetMsgId(uint32)
	//設定訊息的長度
	SetDataLen(uint32)
	//設定訊息的內容
	SetData([]byte)
}
