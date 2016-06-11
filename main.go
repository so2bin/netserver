package main

import (
	"bufio"
	//"common"
	"fmt"
	"netserver/server"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var instr string
	var statelock = false
	var st = time.NewTicker(2 * time.Second)
	svr := new(server.Server)
	go goServerLoop(svr)
	for {
		fmt.Printf("input command>: ")
		scan := bufio.NewScanner(os.Stdin)
		if scan.Scan() {
			instr = scan.Text()
			//fmt.Println(instr)
			switch instr {
			case "exit":
				svr.Destory()
				teardown()
				return
			case "showstate":
				statelock = !statelock
				fmt.Println(statelock)
				if statelock {
					go showState(st.C, svr)
				} else {
					st.Stop()
				}
			default:
				parseCommand(instr)
			}
		} else {
			fmt.Println("wrong  input")
		}
	}
}

// server main loop function
func goServerLoop(svr *server.Server) {
	fmt.Println("server begin listening tcp port....")
	svr.Start()
	defer svr.Destory()
}

func showState(c <-chan time.Time, svr *server.Server) {
	for v := range c {
		v = v
		time.Sleep(time.Second * 2)
		svr.ShowSvrverState()
	}
}

// teardown function when exit
func teardown() {

}

func parseCommand(com string) {

}
