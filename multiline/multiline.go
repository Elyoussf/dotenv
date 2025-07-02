package multiline

import (
	"strings"
)

func ReadSmartly(lines []string) map[string]string {

	// All lines are not comments (verifies from the caller)
	res := make(map[string]string)
	keyAndMultilineVal := ""
	watch := false
	for _, line := range lines {

		if !watch && strings.Contains(line, "=") {
			watch = true
			keyAndMultilineVal += line
		} else if watch && strings.Contains(line, "=") {
			// Now keyAndMultilineVal contains all the required content
			Splitted := strings.SplitN(keyAndMultilineVal, "=", 2)
			key := strings.TrimSpace(Splitted[0])
			val := strings.TrimSpace(Splitted[1])
			res[key] = val
			keyAndMultilineVal = line // Start new Key-Val
		} else if watch && !strings.Contains(line, "=") {
			keyAndMultilineVal += line
		}
	}
	if len(keyAndMultilineVal) != 0 {
		Splitted := strings.SplitN(keyAndMultilineVal, "=", 2)
		key := strings.TrimSpace(Splitted[0])
		val := strings.TrimSpace(Splitted[1])
		res[key] = val
	}
	return res

	// This Code supposes that the rest of Value string does not contain = sign because other wise
	// it would consider it as a new value For example :
	// key=jkjdhjdhjdhjhd
	// dsd=dsdsd  <=== In this case this will be treated as new key with value but i want key=jkjdhjdhjdhjhddsd=dsdsd

}
