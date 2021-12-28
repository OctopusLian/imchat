package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type NetNod struct {
	recvCon   *net.UDPConn
	broadCon  *net.UDPConn
	recvChan  chan []byte
	broadChan chan []byte
	sync.RWMutex
}

var netNod *NetNod = &NetNod{
	recvChan:  make(chan []byte, 1024),
	broadChan: make(chan []byte, 1024),
}

func publish(msg []byte) {

	netNod.broadChan <- msg

}
func StartBroadNode(port int) {

	//初始化发送端口
	broadAddr := &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: port,
	}

	broadCon, err := net.DialUDP("udp", nil, broadAddr)
	defer broadCon.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	netNod.broadCon = broadCon

	for {
		select {
		case d := <-netNod.broadChan:

			broadCon.Write(d)

		}
	}

}

func StartRecvNode(port int) {
	recvAddr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	}
	recvCon, err := net.ListenUDP("udp", recvAddr)
	defer recvCon.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	netNod.recvCon = recvCon
	//初始化变量

	for {
		var buf [250]byte

		n, raddr, err := recvCon.ReadFromUDP(buf[0:])

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if n > 1 && raddr != nil {

		}

		netNod.recvChan <- buf[0:n]
		//判断这个数据是不是自己的,是就处理,不是就抛弃
		//fmt.Println("msg from ",raddr," is ", string(buf[0:n]))
	}
}
func main() {
	go StartRecvNode(3000)
	go StartBroadNode(3000)
	t1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case d := <-netNod.recvChan:
			fmt.Println("<=", string(d))
			break
		case <-t1.C:
			publish([]byte(time.Now().String()))
		}
	}

}
