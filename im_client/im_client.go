package main

import (
	"github.com/golang/glog"
	"sync/atomic"
	"time"
)

var sendNum int64
var receiveNum int64


func main() {

	glog.Infof("begin run im_client \n")

	addr := "127.0.0.1:23000"
	readUid := int64(101000001)
	writeUid := int64(101000002)


	readClient := NewChannel(addr, readUid)
	go readClient.Start()
	writeClient := NewChannel(addr, writeUid)
	go writeClient.Start()

	for   {
		writeClient.sendMsg(101,readUid,"你好 tomorrow better ")
		time.Sleep(5 * time.Second)
	}


	select {}
}

func (client *Channel)sendMsg( seq int, receiver int64, content string) {
	glog.Info("sendMsg --- receiver :",receiver,"   send :",client.uid)
	msg := &Message{MSG_IM, seq, DEFAULT_VERSION, 0, &IMMessage{client.uid, receiver, 0, int32(1), content}}
	client.wt <- msg
	atomic.AddInt64(&sendNum, 1)
	atomic.AddInt64(&client.sendM,1)
}