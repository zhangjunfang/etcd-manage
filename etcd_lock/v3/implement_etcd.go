package v3

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/error"
)

func (lock *EtcdLocker) Acquire(key string, ttl int64) (Lock, error) {

	//uid
	id := addPrefix(lock.Key)
	putRes, err := lock.Client.Put(context.TODO(), id, name)
	if err != nil {
		return nil, error.Error("put data fail")
	}
	mm := clientv3.WithFirstCreate()
	mm = append(mm, clientv3.WithPrefix())
	res, err := lock.Client.Get(context.TODO(), id, mm...)
	if err != nil {
		return nil, error.Error("failed to get the data ")
	} else {
		if string(res.Kvs[0].Key) == key && string(res.Kvs[0].Value) == id {
			//return  得到的数据
			//hookFun  在这里 预留给客户使用
			resp, _ := cli.Grant(context.TODO(), ttl)
			putRes, err := lock.Client.Put(context.TODO(), id, name, clientv3.WithLease(resp.ID))
			return EtcdLock{
				Client: clientv3,
				Key:    id,
				Index:  putRes.Header.Revision,
			}, nil
		} else {
			return nil, error.Error("failed to get the locker ")
		}
	}
	return nil, error.Error("unkown  err")
}
func (lock *EtcdLocker) WaitAcquire(key string, ttl uint64) (Lock, error) {
	//uid
	id := addPrefix(lock.Key)
	putRes, err := lock.Client.Put(context.TODO(), id, name)
	if err != nil {
		return nil, error.Error("put data fail")
	}
	mm := clientv3.WithFirstCreate()
	mm = append(mm, clientv3.WithPrefix())
	hasLock := false
	for !hasLock {
		res, err := lock.Client.Get(context.TODO(), id, mm...)
		if err != nil {
			return nil, error.Error("failed to get the data ")
		} else {
			if string(res.Kvs[0].Key) == key && string(res.Kvs[0].Value) == id {
				//return  得到的数据
				//hookFun  在这里 预留给客户使用
				resp, _ := cli.Grant(context.TODO(), ttl)
				putRes, err := lock.Client.Put(context.TODO(), id, name, clientv3.WithLease(resp.ID))
				hasLock = true
				return EtcdLock{
					Client: clientv3,
					Key:    id,
					Index:  putRes.Header.Revision,
				}, nil
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
	return nil, error.Error("unkown  err")
}
func (lock *EtcdLock) Release() error {
	_, err := lock.Client.Delete(context.Background(), lock.Key)
	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}
