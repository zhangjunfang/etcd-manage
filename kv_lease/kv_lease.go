package kv_lease

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var cli *clientv3.Client
var err error

func init() {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://192.168.1.90:2379"},
		DialTimeout: time.Second * 5,
	})
}
func KV_lease_grant() {
	defer cli.Close()
	//观察数据变动事件
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo", clientv3.WithPrefix())
		for {
			select {
			case v, ok := <-watchChan:
				{
					if ok {
						for i, va := range v.Events {
							fmt.Println(i, string(va.Kv.Key), string(va.Kv.Value), va.IsCreate(), va.IsModify(), va.PrevKv, va.Type)
						}

					}
				}
			}
		}
	}()
	//创建数据过期时间
	resp, _ := cli.Grant(context.TODO(), 2)
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), "bar"+strconv.Itoa(k), clientv3.WithLease(resp.ID))
	}
	//撤销到期的数据
	//cli.Revoke(context.TODO(), resp.ID)
	time.Sleep(time.Second * 12)
	res, _ := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	for _, v := range res.Kvs {
		fmt.Println("--------------------", string(v.Key), string(v.Value))
	}
}
func KV_lease_revoke() {
	defer cli.Close()
	//观察数据变动事件
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo", clientv3.WithPrefix())
		for {
			select {
			case v, ok := <-watchChan:
				{
					if ok {
						for i, va := range v.Events {
							fmt.Println(i, string(va.Kv.Key), string(va.Kv.Value), va.IsCreate(), va.IsModify(), va.PrevKv, va.Type)
						}

					}
				}
			}
		}
	}()
	//创建数据过期时间
	resp, _ := cli.Grant(context.TODO(), 10)
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), "bar"+strconv.Itoa(k), clientv3.WithLease(resp.ID))
	}
	//撤销到期的数据
	cli.Revoke(context.TODO(), resp.ID)
	time.Sleep(time.Second * 12)
	res, _ := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	for _, v := range res.Kvs {
		fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&", string(v.Key), string(v.Value))
	}
	fmt.Println("--------------KV_lease_revoke() ----------------")
}
func KV_lease_keepAlive() {
	defer cli.Close()
	//观察数据变动事件
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo", clientv3.WithPrefix())
		for {
			select {
			case v, ok := <-watchChan:
				{
					if ok {
						for i, va := range v.Events {
							fmt.Println(i, string(va.Kv.Key), string(va.Kv.Value), va.IsCreate(), va.IsModify(), va.PrevKv, va.Type)
						}

					}
				}
			}
		}
	}()
	//创建数据过期时间
	resp, _ := cli.Grant(context.TODO(), 10)
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), "bar"+strconv.Itoa(k), clientv3.WithLease(resp.ID))
	}
	//到期的数据永久存在
	cli.KeepAlive(context.TODO(), resp.ID)
	time.Sleep(time.Second * 12)
	res, _ := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	for _, v := range res.Kvs {
		fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&", string(v.Key), string(v.Value))
	}
	fmt.Println("--------------KV_lease_keepAlive() ----------------")
}
func KV_lease_keepAlive_once() {
	defer cli.Close()
	//观察数据变动事件
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo", clientv3.WithPrefix())
		for {
			select {
			case v, ok := <-watchChan:
				{
					fmt.Println(time.Now().Second())
					if ok {
						for i, va := range v.Events {
							fmt.Println(i, string(va.Kv.Key), string(va.Kv.Value), va.IsCreate(), va.IsModify(), va.PrevKv, va.Type)
						}

					}
				}
			}
		}
	}()
	//创建数据过期时间
	resp, _ := cli.Grant(context.TODO(), 2)
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), "bar"+strconv.Itoa(k), clientv3.WithLease(resp.ID))
	}
	fmt.Println(time.Now().Second())
	//重新开始计算到期时间
	cli.KeepAliveOnce(context.TODO(), resp.ID)
	fmt.Println(time.Now().Second())
	time.Sleep(time.Second * 12)
	res, _ := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	for _, v := range res.Kvs {
		fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&", string(v.Key), string(v.Value))
	}
	fmt.Println("--------------KV_lease_keepAlive_once() ----------------")
}
