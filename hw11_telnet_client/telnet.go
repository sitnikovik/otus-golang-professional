package main

import (
	"io"
	"log"
	"net"
	"time"
)

// receiveBufferSize is the size of the buffer for receiving data.
const receiveBufferSize = 1024

// TelnetClient describes the Telnet client.
type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	conn    io.ReadWriteCloser
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

// NewTelnetClient returns a new TelnetClient.
func NewTelnetClient(
	address string,
	timeout time.Duration,
	in io.ReadCloser,
	out io.Writer,
) TelnetClient {
	// Place your code here.
	// P.S. Author's solution takes no more than 50 lines.
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Fatalf("tcp err: %v", err)
		return nil
	}

	return &telnetClient{
		conn:    conn,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
func (c *telnetClient) Connect() error {
	// Connection is already established in NewTelnetClient
	return nil
}

func (c *telnetClient) Close() error {
	return c.conn.Close()
}

func (c *telnetClient) Send() error {
	bb := make([]byte, receiveBufferSize)
	_, err := c.in.Read(bb)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(bb)
	return err
}

func (c *telnetClient) Receive() error {
	bb := make([]byte, receiveBufferSize)
	n, err := c.conn.Read(bb)
	if err != nil {
		return err
	}

	_, err = c.out.Write(bb[:n])
	return err
}
