package command

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidResponse = errors.New("gonic: invalid or unexpected response")
)

type Started struct {
	Protocol int
	Buffer   int
}

func ParseStarted(raw string) (*Started, error) {
	if !strings.HasPrefix(raw, "STARTED") {
		return nil, ErrInvalidResponse
	}
	args := parseCommand(raw)
	rawProtocol, rawBuffer := args[2], args[3]
	protocol, err := strconv.Atoi(rawProtocol)
	if err != nil {
		return nil, ErrInvalidResponse
	}

	buffer, err := strconv.Atoi(rawBuffer)
	if err != nil {
		return nil, ErrInvalidResponse
	}

	return &Started{
		Protocol: protocol,
		Buffer:   buffer,
	}, nil
}
