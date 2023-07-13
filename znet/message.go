package znet

type Message struct {
	Id      uint32 //訊息的id
	DataLen uint32 //訊息的長度
	Data    []byte //訊息的內容
}

// 創建一個Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 取得訊息的id
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// 取得訊息的長度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// 取得訊息的內容
func (m *Message) GetMsgData() []byte {
	return m.Data
}

// 設定訊息的id
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

// 設定訊息的長度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

// 設定訊息的內容
func (m *Message) SetData(data []byte) {
	m.Data = data
}
