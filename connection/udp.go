package connection

import (
	"fmt"
	"net"
	"time"
)

func BroadcastUdp(port string) {
	fmt.Println("Broadcasting info via UDP...")

	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:"+port)
	if err != nil {
		fmt.Println("There was an error with the resolving the UDP address! ", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("There was an error dialing UDP! ", err)
		return
	}

	defer conn.Close()

	for {
		message := []byte("DISCOVERY: New peer here!")
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error sending the broadcast! ", err)
			return
		}
		fmt.Println("Broadcast sent!")

		// Broadcast every ten seconds
		time.Sleep(10 * time.Second)
	}
}

func ListenUdp(port string, handlePeer func(string)) {
	fmt.Println("Listening for broadcasts...")

	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address!", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error setting up the listener!", err)
		return
	}

	defer conn.Close()

	for {
		// Create an empty byte array
		buffer := make([]byte, 1024)
		n, srcAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP!", err)
			continue
		}
		message := string(buffer[:n])
		fmt.Printf("Received message: %s from %s\n", message, srcAddr.String())

		if message == "DISCOVERY: New peer here!" {
			handlePeer(srcAddr.IP.String())
		}
	}
}
