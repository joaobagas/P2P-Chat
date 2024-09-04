package main

import (
	conn "P2P-Chat/connection"
	"fmt"
)

func main() {
	var choice int

	fmt.Println("Do you want to be the server (1) or the client (2)?")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		conn.BroadcastUdp()
		return
	case 2:
		conn.ListenUdp()
		return
	default:
		fmt.Println("Unsupported option!")
		return
	}

}
