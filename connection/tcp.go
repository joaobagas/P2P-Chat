package connection

import (
	"fmt"
	"net"
)

func StartTCPListener(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error setting up TCP listener! ", err)
		return
	}

	defer listener.Close()
	fmt.Println("TCP server listening on port: ", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println()
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection! ", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Received message! ", message)
	}
}

func ConnectToPeer(ip string, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected to peer", ip)
	for {
		var message string
		fmt.Print("Enter a message: ")
		fmt.Scanln(&message)

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message! ", err)
			return
		}
	}
}
