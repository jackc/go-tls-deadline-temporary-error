package main

import (
	"crypto/tls"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalln("LoadX509KeyPair err:", err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "127.0.0.1:8000", config)
	if err != nil {
		log.Fatalln("Listen err:", err)
	}
	log.Println("Dummy TLS server listening on:", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept err:", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("Read err:", err)
			return
		}
	}
}
