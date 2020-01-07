package tcp

import (
	"bufio"
	"fmt"
	log "github.com/golang/glog"
	"io"
	"net"
)

type TcpConn struct {
	conn *net.TCPConn
	rc   chan []byte
}

//1对1的tcp
func NewServer(port int) *TcpConn {
	listen_addr := fmt.Sprintf("0.0.0.0:%d", port)
	listen, err := net.Listen("tcp", listen_addr)
	if err != nil {
		log.Errorf("listen err:%s", err)
		return nil
	}
	tcp_listener, ok := listen.(*net.TCPListener)
	if !ok {
		log.Error("listen err")
		return nil
	}

	fmt.Println("wait client...")
	conn, err := tcp_listener.AcceptTCP()
	if err != nil {
		log.Errorf("accept err:%s", err)
		return nil
	}

	server := &TcpConn{
		conn: conn,
		rc:   make(chan []byte),
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!")
	go server.tcpPipe()
	return server
}

func NewClient(ip string) *TcpConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", ip)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return nil
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!")

	client := &TcpConn{
		conn: conn,
		rc:   make(chan []byte),
	}

	go client.tcpPipe()
	return client
}

func (tcp *TcpConn) tcpPipe() {

	//获取一个连接的reader读取流
	reader := bufio.NewReader(tcp.conn)
	//接收并返回消息
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		tcp.rc <- []byte(message)
	}
}

func (tcp *TcpConn) Close() {
	tcp.conn.Close()
}

func (tcp *TcpConn) Write(b []byte) {
	tcp.conn.Write(b)
}

func (tcp *TcpConn) Read() []byte {
	return <-tcp.rc
}
