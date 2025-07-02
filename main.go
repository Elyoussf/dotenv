package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Elyoussf/dotenvparser/linehandler"
)

func main() {
	curr_dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error while getting the current dir")
	}
	dotenvpath := curr_dir + "/.env"
	file, err := os.Open(dotenvpath)
	if err != nil {
		log.Fatal("Error Opening The file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		keyval := linehandler.ReturnKeyVal(line)
		if keyval.Key != "" {
			log.Println(keyval)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error Scanning the file!")
	}
}
