package protocol

import (
	"fmt"
	"net"
	"time"
)

func connectAndSendMsg() {
	// 连接到服务端建立的tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
	fmt.Printf("客户端: %T\n", conn)
	if err != nil {
		// 连接的时候出现错误
		fmt.Println("err :", err)
		return
	}
	// 当函数返回的时候关闭连接
	defer conn.Close()
	for {
		time.Sleep(time.Second * 2)
		write, err := conn.Write([]byte("hello world")) // 发送数据
		if err != nil {
			fmt.Println("发送失败：", err)
			return
		}
		fmt.Println("发送：", write)
	}
}
