/*
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello john!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/

package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
	"gameProto"
	"github.com/golang/protobuf/proto"
)

func main() {

	http.Handle("/websocket", websocket.Handler(echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}


func echo(ws *websocket.Conn) {
	var err error
	for {
		var reply []byte

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Printf("receive error===%v\n", err)
			break
		}

		fmt.Println("reply======= ", reply)
		receiveMsg(reply)
	}
}

func receiveMsg(msg []byte) {
	baseMsg := gameProto.BaseMessage{}
	err := proto.Unmarshal(msg, &baseMsg)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	fmt.Println("msg id======", baseMsg.GetMid())
	fmt.Println("msg str =====", baseMsg.GetSerialization())

	//msgType := gameProto.MessageType_GAME_USERINFO
	//fmt.Println("msg type======", msgType)

	switch baseMsg.GetMid() {
	case int32(gameProto.MessageType_UNKNOWN):
		fmt.Println("msg err")
	case int32(gameProto.MessageType_GAME_USERINFO):
		userInfoMsg := gameProto.GameUserInfo{}
		err = proto.Unmarshal(baseMsg.GetSerialization(), &userInfoMsg)

		fmt.Println("uid====", userInfoMsg.GetUid())
		fmt.Println("uname====", userInfoMsg.GetUname())
		fmt.Println("facebook id====", userInfoMsg.GetFacebookid())
		fmt.Println("setFlagCount====", userInfoMsg.GetSetFlagCount())
		fmt.Println("FoundFlagCount====", userInfoMsg.GetFoundFlagCount())
	default:
		fmt.Println("unknow msg")
	}
}
