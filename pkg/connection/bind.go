// Package connection provides functions for managing network connections.
package connection

import (
	"fmt"
	"log"
	"net"
)

// Bind listens on the specified address and handles incoming connections.
func (C Server) Bind(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v\n", err)
	}
	fmt.Printf("Listening on : %s\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			errMsg := fmt.Errorf("accepting connection failed : %v", err)
			// Handle or log the error:
			fmt.Println(errMsg)
			continue
		}

		go C.handleConnection(conn)
	}
}
