package token

import (
	"bufio"
	"chatLoadTest/data"
	"log"
	"os"
	"strings"
)

func ParseTokens(filePath string) data.Tokens {
	log.Println("Parsing tokens ", filePath)
	var tokens data.Tokens
	tokens.Data = make(map[string]string)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error openning config file", err)
	}
	scanner := bufio.NewScanner(f)
	defer f.Close()

	for scanner.Scan() {
		scanned := strings.Split(scanner.Text(), " ; ")
		if len(scanned) == 2 {
			//log.Println(scanned[0], " ", scanned[1])
			tokens.Data[scanned[0]] = scanned[1]
		}

	}
	return tokens
}
