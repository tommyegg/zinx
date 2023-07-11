package ziface

/*
	封包,拆包,模組
	直接面向TCP連線中的數據流，用於處理TCP黏包問題
*/

type IDataPack interface {
	//獲取包的長度方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg Imessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (Imessage, error)
}
