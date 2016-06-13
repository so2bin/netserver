package server

import (
	"fmt"
	"net"
	"netserver/common"
	"time"
)

type Server struct {
	nodes map[string]netnode // ip:port as the key
}

func (svr *Server) Start() {
	svr.nodes = make(map[string]netnode)
	host := ":6543"
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	common.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	common.CheckErr(err)
	defer listener.Close()
	for {
		tcpconn, err := listener.Accept()
		if err != nil {
			continue
		}
		tcprmstr := tcpconn.RemoteAddr().String()
		node := netnode{conn: tcpconn, host: tcprmstr, lasttimerecvdata: time.Now().Unix()}
		svr.nodes[tcprmstr] = node
		go goDearConn(svr.nodes, tcprmstr)
	}
}

func (svr *Server) CloseIP(host string) {
	if c, ok := svr.nodes[host]; ok {
		c.conn.Close()
		delete(svr.nodes, host)
		fmt.Println("OK delete host:", host)
	} else {
		fmt.Println("Warning:  not find a connected host...")
	}
}

func (svr *Server) Destory() {
	for k, _ := range svr.nodes {
		delete(svr.nodes, k)
	}
}

func (svr *Server) ShowSvrverState() {
	fmt.Printf("server connect num: %d\n", len(svr.nodes))
}

func goDearConn(snd map[string]netnode, lstr string) {
	conn := snd[lstr].conn
	defer conn.Close()
	fmt.Println("node %s connected!", lstr)
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connection is break ", lstr)
			delete(snd, lstr)
			return
		}
		fmt.Println(n, string(buf))
		time.Sleep(time.Second)
	}
}
