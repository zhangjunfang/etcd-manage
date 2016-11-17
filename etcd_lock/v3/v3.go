package v3

import "github.com/coreos/etcd/clientv3"

//分布式锁 Etcd结构体
type EtcdLocker struct {
	Client *clientv3.Client
}

//创建分布式锁 结构体
//func NewEtcdLocker(client *clientv3.Client) *Locker {
//	return &EtcdLocker{client: client}
//}

type EtcdLock struct {
	Client *clientv3.Client
	Key    string
	Index  uint64
}
