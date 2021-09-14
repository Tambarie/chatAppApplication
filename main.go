package main

import (
	"log"
	"net"
)

func main()  {
	s := newServer()
	go s.run()
	listener, err := net.Listen("tcp",  "localhost:6000")

	if err != nil{
		log.Fatalf("Unable to start server: %s", err.Error())
		return
	}

	defer listener.Close()
	log.Println("Server has started on port :6000")

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatalf("Server failed to listen: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}