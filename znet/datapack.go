package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/tommyegg/zinx/utils"
	"github.com/tommyegg/zinx/ziface"
)

//封包，拆包的具體模組

type DataPack struct{}

// 拆包封包實例的一個初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 獲取包的長度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//Datalen uint32(4byte) + ID uint32(4byte)
	return 8
}

// 封包方法
func (dp *DataPack) Pack(msg ziface.Imessage) ([]byte, error) {
	//創建一個存放bytes字節的緩衝
	dataBuff := bytes.NewBuffer([]byte{})

	//將dataLen寫進databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//將MsgId 寫進databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//將data數據 寫進databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// 拆包方法 (將包的head訊息讀出來) 之後再根據head的訊息裡的data的長度，再進行一次讀取
func (dp *DataPack) Unpack(binaryData []byte) (ziface.Imessage, error) {
	//創建一個從輸入二進制數據的ioreader
	dataBuff := bytes.NewReader(binaryData)

	//只解壓head訊息，得到datalen和MsgId
	msg := &Message{}

	//讀dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//讀MsgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判斷datalen是否已經超出了允許的最大包長度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv!")
	}

	return msg, nil
}
