package etcd_queue_disrupt

import (
	"context"
	"strconv"

	"github.com/coreos/etcd/clientv3"
)

const (
	pre = "queue"
)

type Queue interface {
	Put(index int64, v interface{}) error
	Take(index uint64) (interface{}, error)
}

//分布式锁 Etcd结构体
type EtcdQueue struct {
	Client *clientv3.Client
}

func (q *EtcdQueue) Put(index int64, v interface{}) error {
	//q.Client.Put(context.TODO(), pre+"/"+strconv.Itoa(int(time.Now().Unix())), v)
	_, err := q.Client.Put(context.TODO(), pre+"/"+strconv.Itoa(int(index)), v)
	return err
}

func (q *EtcdQueue) Take(index uint64) (interface{}, error) {
	res, err := q.Client.Get(context.TODO(), pre+"/"+strconv.Itoa(int(index)))
	if err != nil {
		return nil, err
	} else {
		return res, err
	}
	return nil, nil
}
