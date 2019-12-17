package main

import (
	"CloudStorage/config"
	dblayer "CloudStorage/db"
	"CloudStorage/mq"
	"CloudStorage/store/ceph"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ProcessTransfer : 处理文件转移
func ProcessTransfer(msg []byte) bool {
	log.Println(string(msg))

	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println("111 " + err.Error())
		return false
	}

	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println("222 " + err.Error())
		return false
	}
	defer fin.Close()
	data, _ := ioutil.ReadAll(fin)
	//fmt.Println(len(data))
	err = ceph.PutObject(
		"userfile",
		pubData.DestLocation,
		data)
	if err != nil {
		log.Println("333 " + err.Error())
		return false
	}

	suc := dblayer.UpdateFileLocation(
		pubData.FileHash,
		pubData.DestLocation)
	if !suc {
		return false
	}
	return true
}

func main() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")
	mq.StartConsume(
		config.TransCephQueueName,
		"transfer_ceph",
		ProcessTransfer)
}
