package command

import (
	"fmt"
	"strings"
)

const (
	Delimiter = "\n"
)

type Request struct {
	Command
	Params []string
}

func (r Request) String() string {
	return fmt.Sprintf("%s %s%s", r.Command, strings.Join(r.Params, " "), Delimiter)
}

func NewRequest(command Command, params ...string) *Request {
	return &Request{
		Command: command,
		Params:  params,
	}
}
