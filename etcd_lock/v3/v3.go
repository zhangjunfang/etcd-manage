package v3

import "github.com/coreos/etcd/clientv3"


//分布式锁  获取锁接口方法
type Locker interface {
	Acquire(key string, ttl clientv3.RetryLeaseClient()) (Lock, error)
	WaitAcquire(key string, ttl uint64) (Lock, error)
	Wait(key string) error
}
//分布式锁 Etcd结构体
type EtcdLocker struct {
	Client *clientv3.Client
}
//创建分布式锁 结构体
func NewEtcdLocker(client *clientv3.Client) *Locker {
	return &EtcdLocker{client: client}
}
//分布式锁 锁释放接口
type Lock interface {
	Release() error
}
type EtcdLock struct {
	Client *clientv3.Client
	key    string
	index  uint64
}