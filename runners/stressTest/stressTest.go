package stressTest

import (
	"chatLoadTest/data"
	"chatLoadTest/runners/wsconnect"
	"log"
	"time"
)

func Run(tokens data.Tokens, config data.ParsedConfig) {
	var i int
	done := make(chan int64)
	go func() {
		var alltime int64
		for {
			select {
			case t := <-done:
				alltime = alltime + t
				log.Println("All runners time so far ", alltime, " ms")
			}

		}
	}()
	for uid, token := range tokens.Data {
		i++
		go wsconnect.DialToWs(config.Addr, uid, token, config.Sendmsg, i, done)

		if i%10 == 0 {
			time.Sleep(80 * time.Millisecond)
		}
		if i%4 == 0 {
			time.Sleep(90 * time.Millisecond)
		}
		if i >= config.Connections || i == len(tokens.Data)-1 {
			log.Println("Max connections reached ", i)
			time.Sleep(30 * time.Second)
			return
		}
	}
}
