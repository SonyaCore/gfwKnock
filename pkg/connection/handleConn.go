package connection

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func (C Server) handleConnection(backend net.Conn) {
	var fragmentCounter = 0

	defer backend.Close()
	backend.SetDeadline(time.Now().Add(C.SocketTimeout))

	buff := new(bytes.Buffer)
	HandShake(backend, buff)

	streamConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", C.CloudFlareIP, C.CloudFlarePort))
	if err != nil {
		log.Printf("conn error : %v", err)
		return
	}

	for {
		fragmentCounter++
		fr := int64(C.LFragment)
		num, err := io.CopyN(streamConn, buff, fr)
		log.Printf("seq #%d: size: %d", fragmentCounter, num)

		if err != nil {
			// check if data is empty
			if err == io.EOF {
				break
			}
			fmt.Errorf("error : %s", err.Error())
			return
		}

		time.Sleep(C.FragmentSleep)

		if buff.Len() == 0 {
			break
		}
	}
	log.Println(strings.Repeat("*", 4), "Finish", strings.Repeat("*", 4))

	stream(backend, streamConn)

}
