package wsconnect

import (
	"encoding/json"
	"log"

	"time"

	"github.com/gorilla/websocket"
)

type msgStruct struct {
	MsgType string      `json:"type"`
	Data    interface{} `json:"data"`
}
type authMsgStruct struct {
	User_id string `json:"user_id"`
	Token   string `json:"token"`
}
type joinMsgStruct struct {
	Channel_id string `json:"channel_id"`
	Hidden     int    `json:"hidden"`
	Reload     bool   `json:"reload"`
}
type sendMsgStruct struct {
	Channel_id string `json:"channel_id"`
	Text       string `json:"text"`
	Color      string `json:"color"`
	Icon       string `json:"icon"`
	Role       string `json:"Role"`
	Mobile     int    `json:"Mobile"`
}

func DialToWs(addr string, userid string, token string, sendmsg bool, i int, done chan int64) {
	defer func() {
		//close(done)
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
		}
	}()
	startTime := time.Now().UnixMilli()
	//log.Println("Starting ")
	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		//log.Println("dial:", err)
		log.Println("ERROR IN CONNECTIONS APPEARD IN ", i, " RUNNER")

	}

	defer c.Close()
	//done := make(chan struct{})
	func() {
		//defer close(done)
		defer func() {
			//close(done)
			if r := recover(); r != nil {
				log.Println("Recovered. Error:\n", r)
			}
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)
			hanndleMsg(message, c, userid, token, startTime, sendmsg, done)

		}
	}()

	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
}
func hanndleMsg(msg []byte, c *websocket.Conn, id string, token string, start int64, sendmessage bool, done chan int64) {
	var parsedMsg msgStruct
	var FinishTime int64
	//log.Println(string(msg))
	err := json.Unmarshal(msg, &parsedMsg)
	if err != nil {
		log.Println("Error parsing incoming msg ", err)
	}
	//log.Println(parsedMsg.MsgType)
	switch parsedMsg.MsgType {
	case "welcome":
		//Send auth {"type":"auth","data":{"user_id":someuser,"token":"sometoken"}}
		authMsg := msgStruct{MsgType: "auth", Data: authMsgStruct{User_id: id, Token: token}}
		jsonAuth, err := json.Marshal(authMsg)
		if err != nil {
			log.Println("auth json error ", err)
		}
		//log.Println(string(jsonAuth))
		c.WriteMessage(websocket.TextMessage, jsonAuth)
	case "success_auth":
		//Send join {"type":"join","data":{"channel_id":"10603","hidden":0,"reload":false}}
		joinMsg := msgStruct{MsgType: "join", Data: joinMsgStruct{Channel_id: "10603", Hidden: 0, Reload: false}}
		jsonJoin, err := json.Marshal(joinMsg)
		if err != nil {
			log.Println("joun json error ", err)
		}
		//log.Println(string(jsonJoin))
		c.WriteMessage(websocket.TextMessage, jsonJoin)
		if !sendmessage {
			FinishTime = (time.Now().UnixMilli() - start)
			done <- FinishTime
			log.Println("Auth Finished in ", FinishTime, " ms")
		}
	case "success_join":
		if sendmessage {
			//send msg {"type":"send_message","data":{"channel_id":"10603","text":"1","color":"simple","icon":"none","role":"","mobile":0}}
			time.Sleep(30 * time.Millisecond)
			sendMsg := msgStruct{MsgType: "send_message", Data: sendMsgStruct{Channel_id: "10603", Text: "LoadTesting", Color: "simple", Icon: "none", Role: "", Mobile: 0}}
			jsonSend, err := json.Marshal(sendMsg)
			if err != nil {
				log.Println("send json error ", err)
			}
			//log.Println(string(jsonSend))
			c.WriteMessage(websocket.TextMessage, jsonSend)
			FinishTime = (time.Now().UnixMilli() - start)
			done <- FinishTime
			log.Println("Join Finished in ", FinishTime, " ms")
		}
	case "error":
		log.Println(string(msg))
	default:
		//log.Println("not handled msg type", parsedMsg.MsgType)

	}
}
