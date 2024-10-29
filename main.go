package main

import (
	conn "P2P-Chat/connection"
	"fmt"
)

func main() {
	const udpPort = "3000"
	const tcpPort = "4000"

	go conn.BroadcastUdp(udpPort)

	// Start listening for UDP discovery messages
	go conn.ListenUdp(udpPort, func(peerIP string) {
		fmt.Println("Discovered peer:", peerIP)
		go conn.ConnectToPeer(peerIP, tcpPort)
	})

	// Start listening for incoming TCP connections
	conn.StartTCPListener(tcpPort)
}
