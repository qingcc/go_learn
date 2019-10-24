package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//client1()
	client2()
}

func client1() {
	ip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()

	conn.Write([]byte("hello"))

	for {
		b := make([]byte, 1)

		os.Stdin.Read(b)

		conn.Write(b)

		fmt.Printf("<%s>\n", conn.RemoteAddr())

	}

}

func client2() {
	ip := net.ParseIP("127.0.0.1")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()

	conn.Write([]byte("hello"))

	for {
		b := make([]byte, 1024)

		//os.Stdin.Read(b)

		n, err := conn.Read(b)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("read %s from <%s>\n", b[:n], conn.RemoteAddr())

	}
}
