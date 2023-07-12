package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 負責測試datapack拆包 封包的單元測試
func TestDataPack(t *testing.T) {
	/*
		模擬的伺服器
	*/
	//1 創建socketTcp
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}

	go func() {
		//2.從客戶端讀取數據，拆包處理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}

			go func(conn net.Conn) {
				//處理客戶端的請求
				//----> 拆包的過程 <-----
				dp := NewDataPack()
				for {
					//1.第一次從conn讀，把封包的head讀出來
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head eror")
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}

					if msgHead.GetMsgLen() > 0 {
						//msg 是有資料的，需要進行第二次讀取
						//2.第二次從conn讀，根據head中的datalen 再讀取data內容
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						//根據datalen的長度再次從io stream中讀取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err", err)
							return
						}

						//完整的一個訊息已經讀取完畢
						fmt.Println("----> Recv MsgId:", msg.Id, ",datalen=", msg.DataLen, ",data=", string(msg.Data))

					}

				}

			}(conn)
		}

	}()

	/*
		模擬客戶端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	//創建一個封包對象dp

	dp := NewDataPack()

	//模擬黏包過程，封裝兩個msg一起送出
	//封裝第一個msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}
	//封裝第二個msg2包
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}

	//將兩個包黏在一起
	sendData1 = append(sendData1, sendData2...)

	//一次性送給伺服器端
	conn.Write(sendData1)

	//客戶端阻塞
	select {}
}
