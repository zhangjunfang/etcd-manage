package etcd_queue_disrupt

import (
	"context"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/zhangjunfang/etcd-manage/util"
)

//分布式锁 Etcd结构体
type DefaultQueue struct {
	Client *clientv3.Client
}

func (q *DefaultQueue) Put(_ int64, v interface{}) error {
	//q.Client.Put(context.TODO(), pre+"/"+strconv.Itoa(int(time.Now().Unix())), v)
	_, err := q.Client.Put(context.TODO(), pre+"/"+strconv.Itoa(int(util.Lower16BitPrivateIP()))+strconv.Itoa(int(time.Now().UnixNano())), v)
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
