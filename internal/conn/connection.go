package conn

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	command2 "github.com/Haato3o/gonic/internal/command"
	"net"
	"strings"
	"sync"
)

var (
	ErrContextCancelled = errors.New("conn: context has been cancelled or timed out")
)

type message struct {
	value string
	err   error
}

type Connection struct {
	lock          *sync.Mutex
	reader        *bufio.Reader
	writer        *bufio.Writer
	maxBufferSize int
	conn          net.Conn
	outgoing      chan error
	incoming      chan *message
}

func (c *Connection) SendWithContext(ctx context.Context, command *command2.Request) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	go c.sendCommandWithContext(ctx, command)

	select {
	case <-ctx.Done():
		return ErrContextCancelled

	case err := <-c.outgoing:
		return err
	}
}

func (c *Connection) ReceiveWithContext(ctx context.Context) (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	go c.waitForResponseWithContext(ctx)

	select {
	case <-ctx.Done():
		return "", ErrContextCancelled
	case msg := <-c.incoming:
		return msg.value, msg.err
	}
}

func (c *Connection) Close() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	close(c.incoming)
	close(c.outgoing)
	return c.conn.Close()
}

func (c *Connection) sendCommandWithContext(ctx context.Context, command *command2.Request) {
	if _, err := c.writer.WriteString(command.String()); err != nil {
		c.outgoing <- err
		return
	}

	err := c.writer.Flush()
	if ctx.Err() != nil {
		return
	}

	c.outgoing <- err
}

func (c *Connection) waitForResponseWithContext(ctx context.Context) {
	rawBytes, err := c.reader.ReadBytes('\n')

	if ctx.Err() != nil {
		return
	}

	if err != nil {
		c.incoming <- &message{
			value: "",
			err:   err,
		}
	} else {
		value := string(rawBytes)

		if strings.HasPrefix("ERR", value) {
			errorReason := value[4:]
			err = errors.New(fmt.Sprintf("connection: %s", errorReason))
		} else if strings.HasPrefix("STARTED", value) {
			var response *command2.Started
			response, err = command2.ParseStarted(value)

			c.maxBufferSize = response.Buffer
		}

		c.incoming <- &message{
			value: value,
			err:   err,
		}
	}
}

func New(address string, port int) (*Connection, error) {
	host := fmt.Sprintf("%s:%d", address, port)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return &Connection{
		lock:     &sync.Mutex{},
		reader:   bufio.NewReader(conn),
		writer:   bufio.NewWriter(conn),
		conn:     conn,
		outgoing: make(chan error, 1),
		incoming: make(chan *message, 1),
	}, nil
}
