package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
)

func main() {
	filename := "E:\\WorkSpace\\Go\\src\\go_dev\\day11\\logagent\\conf\\logagent.conf"
	err := loadConf("ini",filename)
	if err != nil {
		fmt.Printf("load conf failed,err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err !=nil{
		fmt.Printf("load logger failed,err:%v\n",err)
		panic("load logger failed")
		return
	}
	logs.Debug("load conf success ,conf:%v",appConfig)

	collectConf,err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err!= nil{
		logs.Error("init etcd failed,err:%v",err)
		return
	}

	err = tailf.InitTail(collectConf,appConfig.chanSize)
	if err!= nil{
		logs.Error("init tail failed,err:%v",err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err!= nil{
		logs.Error("init kafka failed,err:%v",err)
		return
	}


	logs.Debug("initialize success")


	//go func() {
	//	var count int
	//	for{
	//		count++
	//		logs.Debug("test for logger %d",count)
	//		time.Sleep(time.Millisecond*100)
	//	}
	//}()

	err =serverRun()
	if err!=nil{
		logs.Error("serverRun failed,err:%v",err)
		return
	}
	logs.Info("program exited")
}
