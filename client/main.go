package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatalln("Dial err:", err)
	}

	err = conn.SetDeadline(time.Now())
	if err != nil {
		log.Fatalln("conn.SetDeadline(time.Now()) err", err)
	}

	_, err = conn.Write([]byte("should fail"))
	if err == nil {
		log.Fatalln("conn.Write succeeded when should have failed")
	}

	// Clear deadline
	err = conn.SetDeadline(time.Time{})
	if err != nil {
		log.Fatalln("final conn.SetDeadline(time.Time{}) err", err)
	}

	_, err = conn.Write([]byte("This connection is permanently broken"))
	if err != nil {
		fmt.Println("Write err", err)
	}

	ne := err.(net.Error)
	fmt.Println("ne.Temporary() =>", ne.Temporary())

}
