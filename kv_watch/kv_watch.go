package kv_watch

/*
  注意：
      以下方法不可以同时调用.原因：每个方法都有这行代码：defer cli.Close()
*/
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

func Watcher_watch() {
	defer cli.Close()
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo")
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
	cli.Put(context.TODO(), "foo", "123", clientv3.WithLease(1000))
	cli.Put(context.TODO(), "foo", "234")
	cli.Put(context.TODO(), "foo", "345")
	cli.Delete(context.TODO(), "foo")
	time.Sleep(time.Second * 1)
	res, _ := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	for k, v := range res.Kvs {
		fmt.Println(k, string(v.Key), string(v.Value))
	}
}

func Watcher_watch_range() {
	fmt.Println(cli)
	defer cli.Close()
	go func() {
		watchChan := cli.Watch(context.TODO(), "foo", clientv3.WithRange("foo16"))
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
	cli.Delete(context.TODO(), "foo", clientv3.WithPrefix())
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), strconv.Itoa(k))
	}
	cli.Delete(context.TODO(), "foo", clientv3.WithPrefix())
	time.Sleep(time.Second * 1)
}
func Watcher_watch_prefix() {
	defer cli.Close()
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

	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), strconv.Itoa(k))
	}
	cli.Delete(context.TODO(), "foo", clientv3.WithPrefix())
	time.Sleep(time.Second * 1)
}
func Watcher_watch_progressNotify() {
	defer cli.Close()
	go func() {
		watchChan := cli.Watch(context.Background(), "foo", clientv3.WithProgressNotify())
		for {
			select {
			case v, ok := <-watchChan:
				{
					if ok {
						fmt.Println(v.IsProgressNotify())
					}
				}
			}
		}
	}()
	for k, _ := range make([]int, 16) {
		cli.Put(context.TODO(), "foo"+strconv.Itoa(k), strconv.Itoa(k))
	}
}
