package connection

import (
	"fmt"
)

const (
	port     = ":8829"
	hostname = "255.255.255.255"
)

func BroadcastUdp() {
	fmt.Println("Broadcasting info via UDP...")

}

func ListenUdp() {
	fmt.Println("Listening for broadcasts...")

	fmt.Println("Broadcast found...")
}
