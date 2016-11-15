package main

import (
	"fmt"

	"github.com/zhangjunfang/etcd-manage/kv_cluster"

	//"github.com/zhangjunfang/etcd-manage/kv_watch"
	//"github.com/zhangjunfang/etcd-manage/kv_watch"
	//"github.com/zhangjunfang/etcd-manage/v_cluster"
)

func main() {
	//kv_watch.Watcher_watch()
	fmt.Println("----------------------------------")
	//kv_watch.Watcher_watch_range()
	fmt.Println("----------------------------------")
	//kv_watch.Watcher_watch_prefix()
	fmt.Println("----------------------------------")
	//	kv_watch.Watcher_watch_progressNotify()
	fmt.Println("----------------------------------")
	//kv_auth.KV_auth()
	fmt.Println("----------------------------------")
	kv_cluster.KV_cluster()
}
