package main

import (
	conn "P2P-Chat/connection"
	"fmt"
)

func main() {
	var choice int
	var username string

	fmt.Println("Do you want to be the server (1) or the client (2)?")
	fmt.Scan(&choice)

	fmt.Println("What is going to be your username?")
	fmt.Scan(&username)

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
