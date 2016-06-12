package server

import (
	"net"
)

type netnode struct {
	conn             net.Conn
	host             string
	lasttimerecvdata int64 // last time recv data (second)
}
