package connection

import (
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"time"
)

type ReaderConn struct {
	io.Reader
}

func (r *ReaderConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (r *ReaderConn) Close() error {
	return nil
}

func (r *ReaderConn) LocalAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

func (r *ReaderConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

func (r *ReaderConn) SetDeadline(t time.Time) error {
	return nil
}

func (r *ReaderConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (r *ReaderConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func HandShake(conn net.Conn, b *bytes.Buffer) *tls.ClientHelloInfo {
	var tlsEssentials tls.ClientHelloInfo
	var tlsConfig = &tls.Config{
		GetConfigForClient: func(tlsHandshake *tls.ClientHelloInfo) (*tls.Config, error) {
			tlsEssentials = *tlsHandshake
			return nil, nil
		}}

	data := &ReaderConn{io.TeeReader(conn, b)}
	tls.Server(data, tlsConfig).Handshake()

	return &tlsEssentials
}
