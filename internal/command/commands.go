package command

type Command string

const (
	CmdStart   Command = "START"
	CmdQuery   Command = "QUERY"
	CmdSuggest Command = "SUGGEST"
	CmdList    Command = "LIST"
	CmdPing    Command = "PING"
	CmdHelp    Command = "HELP"
	CmdQuit    Command = "QUIT"
	CmdPush    Command = "PUSH"
	CmdPop     Command = "POP"
	CmdCount   Command = "COUNT"
	CmdFlushC  Command = "FLUSHC"
	CmdFlushB  Command = "FLUSHB"
	CmdFlushO  Command = "FLUSHO"
	CmdTrigger Command = "TRIGGER"
	CmdInfo    Command = "INFO"
)
