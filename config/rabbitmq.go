package config

const (
	// AsyncTransferEnable : 是否开启文件异步转移(false是默认同步,true是异步的)
	AsyncTransferEnable = true

	// RabbitURL : rabbitmq服务的入口url
	RabbitURL = "amqp://guest:guest@127.0.0.1:5672/"

	// TransExchangeName : 用于文件transfer的交换机
	TransExchangeName = "uploadserver.trans"

	// TransCephQueueName : ceph转移队列名
	TransCephQueueName = "uploadserver.trans.ceph"

	// TransCephErrQueueName : ceph转移失败后写入另一个队列的队列名
	TransCephErrQueueName = "uploadserver.trans.ceph.err"

	// TransCephRoutingKey : 从ceph存储转移的routingkey
	TransCephRoutingKey = "ceph"
)
