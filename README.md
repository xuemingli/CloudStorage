# CloudStorage
>>基于go语言实现的一种分布式云存储服务。
## 环境
* Linux:
```Bash
lxm@lxm:~$ cat /proc/version
Linux version 5.0.0-37-generic (buildd@lcy01-amd64-023) (gcc version 7.4.0 (Ubuntu 7.4.0-1ubuntu1~18.04.1)) #40~18.04.1-Ubuntu SMP Thu Nov 14 12:06:39 UTC 2019
```
* golang:
```Bash
lxm@lxm:~$ go version
go version go1.13.5 linux/amd64
```
* docker:
```Bash
lxm@lxm:~$ sudo docker version
Client: Docker Engine - Community
 Version:           19.03.5
 API version:       1.40
 Go version:        go1.12.12
 Git commit:        633a0ea838
 Built:             Wed Nov 13 07:29:52 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.5
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.12
  Git commit:       633a0ea838
  Built:            Wed Nov 13 07:28:22 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
```
* mysql
```Bash
lxm@lxm:~$ sudo docker images | grep mysql
mysql                                                                            5.7                 1e4405fe1ea9        4 months ago        437MB
```
* Redis
```Bash
lxm@lxm:~$ redis-server --version
Redis server v=4.0.9 sha=00000000:0 malloc=jemalloc-3.6.0 bits=64 build=9435c3c2879311f3
lxm@lxm:~$ redis-cli --version
redis-cli 4.0.9
```
* RabbitMQ
```Bash
lxm@lxm:~$ sudo docker images | grep rabbit
rabbitmq                                                                         management          6bd1749b8197        3 months ago        181MB
rabbitmq                                                                         latest              458123c67b79        3 months ago        151MB
```
* Ceph
```Bash
lxm@lxm:~$ sudo docker images | grep ceph
ceph/radosgw                                                                     latest              6c0c6063916a        4 years ago         594MB
moxiaomomo/ceph-radosgw                                                          latest              6c0c6063916a        4 years ago         594MB
ceph/mon                                                                         latest              b18ef381c929        4 years ago         594MB
moxiaomomo/ceph-mon                                                              latest              b18ef381c929        4 years ago         594MB
ceph/osd                                                                         latest              15adde780f94        4 years ago         596MB
moxiaomomo/ceph-osd                                                              latest              15adde780f94        4 years ago         596MB
```
* consul
```Bash
lxm@lxm:~$ sudo docker images | grep consul
consul                                                                           latest              61c55d0793c6        4 months ago        117MB
```
* kubernetes
```Bash
lxm@lxm:~$ sudo docker images | grep kube
registry.aliyuncs.com/google_containers/kube-proxy                               v1.14.1             20a2d7035165        12 months ago       82.1MB
registry.aliyuncs.com/google_containers/kube-apiserver                           v1.14.1             cfaa4ad74c37        12 months ago       210MB
registry.aliyuncs.com/google_containers/kube-controller-manager                  v1.14.1             efb3887b411d        12 months ago       158MB
registry.aliyuncs.com/google_containers/kube-scheduler                           v1.14.1             8931473d5bdb        12 months ago       81.6MB
k8s.gcr.io/kubernetes-dashboard-amd64                                            v1.10.1             f9aed6605b81        15 months ago       122MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64   v1.10.1             f9aed6605b81        15 months ago       122MB
```
## go依赖包
```Bash
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
#go get github.com/garyburd/redigo/redis
go get github.com/gomodule/redigo/redis
go get github.com/json-iterator/go
go get github.com/aliyun/aliyun-oss-go-sdk/oss
go get gopkg.in/amz.v1/aws
go get gopkg.in/amz.v1/s3
go get github.com/streadway/amqp
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/micro/go-micro
go get github.com/mitchellh/mapstructure
go get github.com/jteeuwen/go-bindata/...
go get github.com/moxiaomomo/go-bindata-assetfs/...
go get github.com/gin-gonic/contrib/static
go get github.com/micro/go-micro/cmd
go get github.com/micro/go-plugins/registry/kubernetes
go get -v github.com/micro/go-plugins/wrapper/breaker/hystrix
go get -v github.com/juju/ratelimit
go get -v github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit
```
## 运行
* 微服务架构下启动方式(非容器化部署):
 * 一键启动微服务(start-all.sh):
 ```Bash
 > cd $GOPATH/src/filestore-server
 > ./service/start-all.sh 
 编译完成:  service/bin/dbproxy
 编译完成:  service/bin/upload
 编译完成:  service/bin/download
 编译完成:  service/bin/transfer
 编译完成:  service/bin/account
 编译完成:  service/bin/apigw
 已启动  dbproxy
 已启动  upload
 已启动  download
 已启动  transfer
 已启动  account
 已启动  apigw
 微服务启动完毕.
 ```
 * 一键关闭微服务(stop-all.sh):
 ```Bash
 > cd $GOPATH/src/filestore-server
 > ./service/stop-all.sh 
 已关闭:  apigw
 已关闭:  account
 已关闭:  transfer
 已关闭:  download
 已关闭:  upload
 已关闭:  dbproxy
 执行完毕.
 ```
