package main

import (
	"bufio"
	"fmt"
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
	port    int
	timeout time.Duration

	stdin  io.Reader
	stdout io.Writer
}

// NewTelnetClient creates a new TelnetClient.
func NewTelnetClient(addr string, port int, timeout time.Duration, stdin io.Reader, stdout io.Writer) TelnetClient {
	return &telnetClient{
		addr:    addr,
		port:    port,
		timeout: timeout,
		stdin:   stdin,
		stdout:  stdout,
	}
}

// Connect connects to the server.
func (t *telnetClient) Connect() error {
	var err error
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.addr, t.port), t.timeout)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}

// Send sends data to the server.
func (t *telnetClient) Send() error {
	reader := bufio.NewReader(t.stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = t.conn.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

// Receive receives data from the server.
func (t *telnetClient) Receive() error {
	reader := bufio.NewReader(t.conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(t.stdout, message)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the server.
func (t *telnetClient) Close() error {
	if t.conn != nil {
		return t.conn.Close()
	}
	return nil
}
