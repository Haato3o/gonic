package command

import "strings"

func parseCommand(rawCommand string) []string {
	fields := strings.Fields(rawCommand)

	// Parse functions
	for idx, field := range fields {
		if !strings.HasSuffix(field, ")") {
			continue
		}

		// TODO: Implement parsing for multiple arguments if needed
		fields[idx] = strings.FieldsFunc(field, func(r rune) bool {
			return r == '(' || r == ')'
		})[1]
	}

	return fields
}
