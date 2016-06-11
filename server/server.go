package server

import (
	"fmt"
	"net"
	"netserver/common"
)

type Server struct {
	conns map[string]net.Conn // ip:port as the key
}

func (svr *Server) Start() {
	host := ":6543"
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	common.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	common.CheckErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go goDearConn(conn)
	}
}

func (svr *Server) CloseIP(host string) {
	if c, ok := svr.conns[host]; ok {
		c.Close()
		delete(svr.conns, host)
		fmt.Println("OK delete host:", host)
	} else {
		fmt.Println("Warning:  not find a connected host...")
	}
}

func (svr *Server) Destory() {
	for k, _ := range svr.conns {
		delete(svr.conns, k)
	}
}

func (svr *Server) ShowSvrverState() {
	fmt.Printf("server connect num: %d\n", len(svr.conns))
}

func goDearConn(conn net.Conn) {
	defer conn.Close()
	for {

	}
}
