package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go_dev/day11/logagent/tailf"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []tailf.CollectConf, err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, localIp := range localIpArray {
		etcdKey := fmt.Sprintf("%s%s", key, localIp)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed,err:%v", err)
			continue
		}
		cancel()

		for _, v := range resp.Kvs {
			fmt.Printf("v:%v", v)
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed,err:%v", err)
				}
				logs.Debug("log config is %v", collectConf)
			}
		}
	}

	initEtcdWatcher()
	return
}

func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

func watchKey(key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key[%s],Unmarshal[%s],err:%v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd ,%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				tailf.UpdateConfig(collectConf)
			}
		}

	}
}
