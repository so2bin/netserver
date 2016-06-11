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
	go goServerLoop()
	for {
		fmt.Printf("input command>: ")
		scan := bufio.NewScanner(os.Stdin)
		if scan.Scan() {
			instr = scan.Text()
			fmt.Println(instr)
			switch instr {
			case "exit":
				teardown()
				return
			case "showstate":
				statelock = !statelock
				if statelock {
					fmt.Println(statelock)
					go showState(st.C)
				} else {
					fmt.Println(statelock)
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
func goServerLoop() {
	svr := new(Server)
	svr.Start()
	defer svr.Destory()
}

func showState(c <-chan time.Time) {
	for v := range c {
		time.Sleep(time.Second * 2)
		fmt.Println(v, "showstate...")
	}
}

// teardown function when exit
func teardown() {

}

func parseCommand(com string) {

}
