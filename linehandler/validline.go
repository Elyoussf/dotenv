package linehandler

// Suppose a value is litterally fitting one line
import (
	"strings"
)

type KeyVal struct {
	Key string
	Val string
}

func ReturnKeyVal(line string) KeyVal {
	if !strings.Contains(line, "=") || (len(line) > 0 && strings.TrimSpace(string(line[0])) == "=") {
		return KeyVal{}
	}
	parts := strings.SplitN(line, "=", 2)

	if len(parts) == 1 {
		return KeyVal{Key: parts[0], Val: ""}
	}
	return KeyVal{Key: strings.TrimSpace(parts[0]), Val: strings.TrimSpace(parts[1])}
}
