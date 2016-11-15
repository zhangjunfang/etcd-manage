package kv_cluster

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var cli *clientv3.Client
var err error

func init() {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:        []string{"http://192.168.1.90:2379"},
		DialTimeout:      time.Second * 5,
		AutoSyncInterval: time.Second * 0,
		Username:         "root",
		Password:         "root",
	})
}
func KV_cluster() {
	defer cli.Close()
	fmt.Println(cli)
	//cli.MemberAdd(context.Background(), cli.Endpoints())
	list, err := cli.MemberList(context.TODO())
	if err == nil {
		for k, v := range list.Members {
			fmt.Println(
				k,
				v.PeerURLs,
			)
			cli.MemberUpdate(context.TODO(), v.ID, cli.Endpoints())
			cli.MemberAdd(context.Background(), cli.Endpoints())
			//cli.MemberRemove(context.Background(), v.ID)
		}
	}
	list, _ = cli.MemberList(context.TODO())
	fmt.Println(
		len(list.Members),
	)
	//cli.MemberAdd(context.Background(), cli.Endpoints())
}
