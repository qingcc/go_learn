package main

import (
	"fmt"
	"net"
)

func main() {
	udp()
}

func udp() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)

		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
