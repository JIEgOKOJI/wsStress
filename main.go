// chatLoadTest project main.go
package main

import (
	"chatLoadTest/data"
	"chatLoadTest/handlers/config"
	"chatLoadTest/handlers/token"
	"chatLoadTest/runners/stressTest"

	//	"chatLoadTest/runners/wsconnect"
	"fmt"
	"log"
)

var Configuration data.ParsedConfig
var ParsedTokens data.Tokens

func main() {
	fmt.Println("Hello World!")
	config := config.Config{Path: "config.yaml"}
	Configuration = config.Load()
	log.Println(Configuration)
	ParsedTokens := token.ParseTokens(Configuration.Tokens)
	log.Println(ParsedTokens)
	stressTest.Run(ParsedTokens, Configuration)
	//wsconnect.DialToWs(Configuration.Addr)
	//config.Config.Load("config.yaml")
}
