package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Elyoussf/dotenvparser/dotenvlocator"
	"github.com/Elyoussf/dotenvparser/linehandler"
	"github.com/Elyoussf/dotenvparser/multiline"
)

func main() {
	// if not specified it is not a multiline
	// if it is and it is true it will be consider it multiline
	// Otherwise willnot  be multiline
	MultilineMode := false
	if len(os.Args) >= 2 {
		MultilineMode = strings.EqualFold(strings.ToUpper(os.Args[1]), strings.ToUpper("true"))
	}
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	dotenvPath, err := dotenvlocator.LocateDotEnv(currDir, ".env")
	if err != nil {
		fmt.Println(err)

	} else {
		file, err := os.Open(dotenvPath)
		if err != nil {
			log.Fatal("Error opening the .env file:", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		result := make(map[string]string)
		lines := []string{}
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			index := strings.IndexRune(line, '#')
			if index != -1 {
				line = line[:index]
			}
			if len(line) == 0 || string(line[0]) == "#" {
				continue
			}

			if MultilineMode {
				lines = append(lines, line)
				continue
			}
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			kv := linehandler.ReturnKeyVal(line)
			if kv.Key != "" {
				if kv.Key[0] == kv.Key[len(kv.Key)-1] && (string(kv.Key[0]) == "\"" || string(kv.Key[0]) == "'") {
					kv.Key = kv.Key[1 : len(kv.Key)-1]
				}
				if kv.Val[0] == kv.Val[len(kv.Val)-1] && (string(kv.Val[0]) == "\"" || string(kv.Key[0]) == "'") {
					kv.Val = kv.Val[1 : len(kv.Val)-1]
				}
				result[kv.Key] = kv.Val
			}
		}
		if MultilineMode {
			result = multiline.ReadSmartly(lines)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Error scanning the file:", err)
		}
		// At this point i have a ready to output map that holds
		// Key - val
		// This where some values processing is done

		jsonBytes, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal("Error encoding to JSON:", err)
		}

		fmt.Println(string(jsonBytes))
	}
}
