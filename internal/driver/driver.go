package driver

import (
	"context"
	command2 "github.com/Haato3o/gonic/internal/command"
	"github.com/Haato3o/gonic/internal/conn"
)

type Mode string

const (
	ModeSearch  Mode = "search"
	ModeIngest  Mode = "ingest"
	ModeControl Mode = "control"
)

type Driver struct {
	mode Mode
	*conn.Connection
}

func (d Driver) ConnectWithContext(ctx context.Context, password string) error {
	// Skip connected message
	if _, err := d.ReceiveWithContext(ctx); err != nil {
		return err
	}

	cmd := command2.NewRequest(command2.CmdStart, string(d.mode), password)
	if err := d.SendWithContext(ctx, cmd); err != nil {
		return err
	}

	_, err := d.ReceiveWithContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func New(address string, port int, mode Mode) (*Driver, error) {
	connection, err := conn.New(
		address,
		port,
	)

	if err != nil {
		return nil, err
	}

	return &Driver{
		mode:       mode,
		Connection: connection,
	}, nil
}
