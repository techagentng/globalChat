package  main

import (
	"log"
	"net"
)

func main()  {
	s := newServer() //Init main server for managing room state
	go s.run()
	//Start TCP server
	lister, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer lister.Close()
	log.Printf("Server started at :8888")
	//Start TCP server END

	for {
		conn, err := lister.Accept() //conn refers to the current client
		if err != nil{
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}

}