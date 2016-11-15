package kv_base

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

var cli *clientv3.Client
var err error

func init() {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://192.168.1.90:2379"},
		DialTimeout: time.Second * 5,
	})
}

//添加key value  c一定不可以为空
func PutKeyValue(c context.Context, key, value string) error {
	if err != nil {
		return err
	}
	cxt, canc := context.WithTimeout(c, time.Second*5)
	res, err := cli.Put(cxt, key, value)
	canc()
	fmt.Println(res.Header.Revision, "-------------------", res.Header.String())
	return err
}

// 错误处理
//通过测试发现  ：  Put时，必须填写key值  value值不作为判断条件
func PutKV_ErrorHandle(c context.Context, key, value string) error {
	if err != nil {
		return err
	}
	cxt, canc := context.WithTimeout(c, time.Second*5)
	res, err := cli.Put(cxt, key, value)
	canc()
	if err != nil {
		switch err {
		case context.Canceled:
			{
				fmt.Println("context.Canceled :", context.Canceled)
			}
		case context.DeadlineExceeded:
			{
				fmt.Println("context.DeadlineExceeded:", context.DeadlineExceeded)
			}
		case rpctypes.ErrEmptyKey:
			{
				fmt.Printf("client-side error: %v\n", err)
			}
		case rpctypes.ErrGRPCTimeout:
			{
				fmt.Printf("client-side error: %v\n", err)
			}
		default:
			{
				fmt.Printf("client-side error: %v\n", err)
			}
		}
	} else {
		fmt.Println(res.Header.Revision, "-------------------", res.Header.String())
	}
	return err
}

//获取key对应的值
func GetKV(c context.Context, key string) {
	if err != nil {
		return
	}
	PutKeyValue(c, "foo", "2")
	ctx, canc := context.WithTimeout(c, time.Second*5)
	res, err := cli.Get(ctx, key)
	canc()
	if err != nil {
		return
	}
	for k, v := range res.Kvs {
		fmt.Println(k, "---", string(v.Key), "----", string(v.Value), "----", v.Version)
	}
}

//获取key对应的value值  带着版本号
func GetKV_WithRev(c context.Context, key string) {
	if err != nil {
		return
	}
	var i int = 0
	i++
	ctx, canc := context.WithTimeout(c, time.Second*5)
	res, _ := cli.Put(ctx, "foo", "foo"+strconv.Itoa(i))
	i++
	ctx, canc = context.WithTimeout(c, time.Second*5)
	cli.Put(ctx, "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	i++
	cli.Put(context.TODO(), "foo", "foo"+strconv.Itoa(i))
	ctx, canc = context.WithTimeout(c, time.Second*5)
	r, _ := cli.Get(ctx, "foo", clientv3.WithRev(res.Header.Revision))
	canc()
	for k, v := range r.Kvs {
		fmt.Println(k, "---", string(v.Key), "----", string(v.Value), "----", v.Version)
	}
	for i := int64(0); i < r.Count; i++ {
		fmt.Println(i, string(r.Kvs[i].Key))
	}
}
func GetKV_WithSortedPrefix() {
	if err != nil {
		return
	}

	for v := range make([]int, 8) {
		ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
		cli.Put(ctx, "foo"+strconv.Itoa(v), strconv.Itoa(v))
		canc()
	}
	ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
	res, _ := cli.Get(ctx, "foo", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	canc()
	for k, v := range res.Kvs {
		fmt.Println(k, string(v.Key))
	}

}
func DeleteKV() {
	if err != nil {
		return
	}
	ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
	defer canc()
	gresp, _ := cli.Get(ctx, "foo", clientv3.WithPrefix())
	fmt.Println(gresp.Count)
	cli.Delete(ctx, "foo", clientv3.WithPrefix())
	gresp, _ = cli.Get(ctx, "foo", clientv3.WithPrefix())
	fmt.Println(gresp.Count)

}
func Compact_rev_KV() {
	if err != nil {
		return
	}
	ctx, canc1 := context.WithTimeout(context.Background(), time.Second*5)
	gresp, _ := cli.Get(ctx, "foo1")
	defer canc1()

	ctx, canc2 := context.WithTimeout(context.Background(), time.Second*5)
	res, _ :=
		cli.Compact(ctx, gresp.Header.Revision)
	defer canc2()

	fmt.Println(res.Header.Revision, gresp.Header.Revision)

}
func TXN_KV() {
	kv := clientv3.NewKV(cli)
	kv.Put(context.TODO(), "key", "xyz")
	ctx, canc2 := context.WithTimeout(context.Background(), time.Second*5)
	defer canc2()
	kv.Txn(ctx).
		If(clientv3.
			Compare(clientv3.
				Value("key"), ">", "abc")).
		Then(clientv3.OpPut("key", "XYZ")).
		Else(clientv3.OpPut("key", "abc")).
		Commit()
	gresp, _ := kv.Get(context.TODO(), "key")
	fmt.Println(string(gresp.Kvs[0].Value))
}
func ExampleKV_do() {

	if err != nil {
		return
	}

	ops := []clientv3.Op{
		clientv3.OpPut("put-key", "123"),
		clientv3.OpGet("put-key"),
		clientv3.OpPut("put-key", "456")}
	for _, op := range ops {
		if _, err := cli.Do(context.TODO(), op); err != nil {
			log.Fatal(err)
		}
	}

}

//func main() {
//defer cli.Close()
//c := context.Background()
//	PutKeyValue(c, "foo", "1")
//	time.Sleep(time.Second * 3)
//	GetKV(c, "foo")
//GetKV_WithRev(c, "foo")
//GetKV_WithSortedPrefix()
//DeleteKV()
//Compact_rev_KV()
//TXN_KV()
//ExampleKV_do()
//}
