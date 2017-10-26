package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func getStatus(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		info := `{"uid": 1, "部门": ["财务", "办公室"]， "好友": ["李四"]}`

		msg := "收到: " + reply + "\n返回: " + info

		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/login", websocket.Handler(getStatus))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
