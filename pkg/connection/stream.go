package connection

import (
	"io"
	"net"
	"sync"
)

func stream(backend net.Conn, client net.Conn) {
	var wg sync.WaitGroup
	wg.Add(2)

	// run the backend (cloudflare)
	go func() {
		io.Copy(backend, client)
		backend.Close()
		defer wg.Done()

	}()

	// run the client (localhost)
	go func() {
		io.Copy(client, backend)
		client.Close()
		defer wg.Done()
	}()
	wg.Wait()
}
