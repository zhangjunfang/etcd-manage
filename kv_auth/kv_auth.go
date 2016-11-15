package kv_auth

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
		Endpoints:   []string{"http://192.168.1.90:2379"},
		DialTimeout: time.Second * 5,
	})
}
func KV_auth() {
	defer cli.Close()
	cli.AuthEnable(context.TODO())
	//添加角色
	role, err := cli.RoleAdd(context.TODO(), "role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(role)
	}
	user, err := cli.UserAdd(context.TODO(), "root", "root")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(user)
	}
	grantRole, err := cli.UserGrantRole(context.TODO(), "root", "role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(grantRole)
	}
	cli.RoleGrantPermission(context.TODO(), "r", "foo", "foo15", clientv3.PermissionType(clientv3.PermReadWrite))
}
