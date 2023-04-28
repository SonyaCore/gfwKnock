package connection

import (
	"fmt"
	"log"
	"net"
)

func (C Server) Bind(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v\n", err)
	}
	fmt.Printf("Listening on : %s\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Errorf("accepting connection failed : %v", err)
			continue
		}

		go C.handleConnection(conn)
	}
}
