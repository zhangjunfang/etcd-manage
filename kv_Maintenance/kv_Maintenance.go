package kv_Maintenance

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
func KV_Maintenance_status() {
	defer cli.Close()
	main := clientv3.NewMaintenance(cli)
	status, _ := main.Status(context.TODO(), cli.Endpoints()[0])
	fmt.Println(status.DbSize)
	fmt.Printf("endpoint: %s / IsLeader: %v\n", cli.Endpoints()[0], status.Header.MemberId == status.Leader)
}

//获取集群中所有的告警信息
func KV_Maintenance_alarmlist() {
	defer cli.Close()
	main := clientv3.NewMaintenance(cli)
	res, _ := main.AlarmList(context.TODO())
	for _, v := range res.Alarms {
		fmt.Println(v.Alarm, v.MemberID)
	}

}

//解除告警信息
func KV_Maintenance_alarmDisarm() {
	defer cli.Close()
	//	main := clientv3.NewMaintenance(cli)
	//	res, _ := main.AlarmList(context.TODO())
	//	for _, m := range cli.Endpoints() {
	//		status, _ := main.Status(context.TODO(), m)
	//		for _, v := range res.Alarms {
	//			fmt.Println(v.Alarm, v.MemberID)
	//			if status.Header.MemberId == v.MemberID {
	//				main.AlarmDisarm(context.TODO(), v)
	//			}

	//		}
	//	}

}

//资源回收 或者 资源释放  针对每一个集群成员进行资源释放
func KV_Maintenance_defragment() {
	defer cli.Close()
	for _, v := range cli.Endpoints() {
		cli.Defragment(context.TODO(), v)
	}
}

func KV_Maintenance_snapshot() {
	defer cli.Close()
	main := clientv3.NewMaintenance(cli)
	reader, _ := main.Snapshot(context.TODO())
	buffer := make([]byte, 512)
	for {
		reader.Read(buffer)

		fmt.Print(string(buffer))
	}
}
