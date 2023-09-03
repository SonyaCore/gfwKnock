package connection

import (
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"time"
)

// ReaderConn is a custom net.Conn implementation that acts as an io.Reader.
type ReaderConn struct {
	io.Reader
}

// Write implements the Write method of the net.Conn interface.
func (r *ReaderConn) Write(b []byte) (int, error) {
	return len(b), nil
}

// Close implements the Close method of the net.Conn interface.
func (r *ReaderConn) Close() error {
	return nil
}

// LocalAddr returns a local network address.
func (r *ReaderConn) LocalAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

// RemoteAddr returns a remote network address.
func (r *ReaderConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

// SetDeadline implements the SetDeadline method of the net.Conn interface.
func (r *ReaderConn) SetDeadline(_ time.Time) error {
	return nil
}

// SetReadDeadline implements the SetReadDeadline method of the net.Conn interface.
func (r *ReaderConn) SetReadDeadline(_ time.Time) error {
	return nil
}

// SetWriteDeadline implements the SetWriteDeadline method of the net.Conn interface.
func (r *ReaderConn) SetWriteDeadline(_ time.Time) error {
	return nil
}

// HandShake performs a TLS handshake on the provided connection and returns
// information about the client's hello message.
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
