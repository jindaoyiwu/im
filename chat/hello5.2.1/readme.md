package ctrl

import (
	"net/http"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"strconv"
	"log"
	"sync"
	"fmt"
)
//本核心在于形成userid和Node的映射关系
//
//定义节点
type Node struct {
	Conn *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}
var clientMap map[int64]*Node = make(map[int64]*Node,0)
var rwlocker sync.RWMutex
//
// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(writer http.ResponseWriter,
	request *http.Request) {

	//todo 检验接入是否合法
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId ,_ := strconv.ParseInt(id,10,64)
	isvalid := checkToken(userId,token)

	//todo 获得conn

	conn,err:=(&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer,request,nil)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	node := &Node{
		Conn:conn,
		DataQueue:make(chan []byte,50),
		GroupSets:set.New(set.ThreadSafe),
	}
	//todo userid和node形成绑定关系
	//map 的并发
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	//todo 完成发送逻辑,con
	go sendproc(node)
	//todo 完成接收逻辑
	go recvproc(node)

	sendMsg(userId,[]byte("hello,world!"))
}

//发送协程
func sendproc(node *Node) {
	for{
		select {
		case data := <-node.DataQueue:
			err :=node.Conn.WriteMessage(websocket.TextMessage,data)
			if err!=nil{
				log.Println(err.Error())
				return
			}
		}
	}
}
//接收协程
func recvproc(node *Node) {
	for{
		_,data,err:=node.Conn.ReadMessage()
		if err!=nil{
			log.Println(err.Error())
			return
		}
		//todo data做进一步的处理
		fmt.Printf("recv %s",data)
		//node.DataQueue<-[]byte(fmt.Sprintf("recv<=%s",data))
	}
}

//发送消息
func sendMsg(userid int64,msg []byte) {
	rwlocker.RLock()
	node,ok:= clientMap[userid]
	rwlocker.RUnlock()

	if ok{
		node.DataQueue<-msg
	}

}
//检测是否有效
func checkToken(userId int64,token string)bool{
	return true
}