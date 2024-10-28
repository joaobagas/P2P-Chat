package connection

import (
	"fmt"
	"net"
)

const (
	port = ":8829"
	ip = "255.255.255.255"
)

func BroadcastUdp() {
	fmt.Println("Broadcasting info via UDP...")

	connection, err := net.ListenPacket("udp4", port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	address, err := net.ResolveUDPAddr("udp4", ip + port)
	if err != nil {
		panic(err)
	}

	// Here I should send the data.
	_, err = connection.WriteTo([]byte("data to transmit"), address)
	if err != nil {
		panic(err)
	}
}

func ListenUdp() {
	fmt.Println("Listening for broadcasts...")

	connection, err := net.ListenPacket("udp4", port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	buffer := make([]byte, 1024)
	n, address, err := connection.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s sent this: %s\n", address, buffer[:n])
}
