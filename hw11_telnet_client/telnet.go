package main

import (
	"io"
	"net"
	"time"
)

// TelnetClient describes the Telnet client.
type TelnetClient interface {
	io.Closer

	// Connect connects to the server.
	Connect() error
	// Send sends data to the server.
	Send() error
	// Receive receives data from the server.
	Receive() error
}

// telnetClient implements the TelnetClient interface.
type telnetClient struct {
	conn net.Conn

	addr    string
	timeout time.Duration

	stdin  io.Reader
	stdout io.Writer
}

// NewTelnetClient creates a new TelnetClient.
func NewTelnetClient(
	addr string,
	timeout time.Duration,
	stdin io.Reader,
	stdout io.Writer,
) TelnetClient {
	return &telnetClient{
		addr:    addr,
		timeout: timeout,
		stdin:   stdin,
		stdout:  stdout,
	}
}

// Connect connects to the server.
func (t *telnetClient) Connect() error {
	var err error
	conn, err := net.DialTimeout("tcp", t.addr, t.timeout)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

// Send sends data to the server.
func (t *telnetClient) Send() error {
	_, err := io.Copy(t.conn, t.stdin)
	return err
}

// Receive receives data from the server.
func (t *telnetClient) Receive() error {
	_, err := io.Copy(t.stdout, t.conn)
	return err
}

// Close closes the connection to the server.
func (t *telnetClient) Close() error {
	if t.conn != nil {
		return t.conn.Close()
	}
	return nil
}
