package main

import (
	"fmt"
	"net"
	"log"
)

func main () {

	netListen,err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error: %s", err.Error())
	}

	defer netListen.Close()
	log.Print("server listen wating...")

	for {
		conn,err := netListen.Accept()

		if err != nil {
			continue
		}

		log.Print(conn.RemoteAddr().String(), "tcp connect success")

		handleConnenction(conn)
	}
}

func handleConnenction (conn net.Conn) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil  {
			log.Print(conn.RemoteAddr().String(), "connect error: ", err)
			return
		}

		log.Print(conn.RemoteAddr().String(), string(buffer[:n] ))
	}
}