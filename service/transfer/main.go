package main

import (
	"CloudStorage/config"
	dblayer "CloudStorage/db"
	"CloudStorage/mq"
	"CloudStorage/store/ceph"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/micro/go-micro"
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

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
	)
	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTranserService() {
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

func main() {
	// 文件转移服务
	go startTranserService()

	// rpc 服务
	startRpcService()
}

/*
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
*/
