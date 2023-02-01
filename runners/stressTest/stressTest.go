package stressTest

import (
	"chatLoadTest/data"
	"chatLoadTest/runners/wsconnect"
	"log"
	"time"
)

func Run(tokens data.Tokens, config data.ParsedConfig) {
	var i int
	for uid, token := range tokens.Data {
		i++
		go wsconnect.DialToWs(config.Addr, uid, token, config.Sendmsg)

		if i%10 == 0 {
			time.Sleep(150 * time.Millisecond)
		}
		if i%4 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
		if i >= config.Connections || i == len(tokens.Data) {
			log.Println("Max connections reached ", i)
			time.Sleep(30 * time.Second)
			return
		}
	}
}
