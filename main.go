package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("etcd boot success")
	defer func(cli *clientV3.Client) {
		err := cli.Close()
		if err != nil {
			log.Error(err)
		}
	}(cli)
	for {
		err := NewLeasesLock(cli, "node1")
		if err != nil {
			return
		}
	}
}

// NewLeasesLock new release
func NewLeasesLock(client *clientV3.Client, ip string) error {
	// 创建新的租约
	lease := clientV3.NewLease(client)
	leaseGrantResponse, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
		return err
	}
	leaseId := leaseGrantResponse.ID
	ctx, cancelFunc := context.WithCancel(context.TODO())
	defer cancelFunc()
	// 撤销租约
	defer func(lease clientV3.Lease, ctx context.Context, id clientV3.LeaseID) {
		_, err := lease.Revoke(ctx, id)
		if err != nil {
			log.Error(err)
		}
	}(lease, context.TODO(), leaseId)
	leaseKeepAliveChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	kv := clientV3.NewKV(client)
	txn := kv.Txn(context.TODO())
	txn.If(clientV3.Compare(clientV3.CreateRevision("/dev/lock"), "=", 0)).Then(
		clientV3.OpPut("/dev/lock", ip, clientV3.WithLease(leaseId))).Else(
		clientV3.OpGet("/dev/lock"))
	txnResponse, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if txnResponse.Succeeded {
		fmt.Println("抢到锁了")
		fmt.Printf("选定主节点 %s\n", ip)
		for {
			select {
			case leaseKeepAliveResponse := <-leaseKeepAliveChan:
				if leaseKeepAliveResponse != nil {
					fmt.Println("续租成功,leaseID :", leaseKeepAliveResponse.ID)
				} else {
					fmt.Println("续租失败")
				}
			}
		}
	} else {
		fmt.Println("没抢到锁", txnResponse.Responses[0].GetResponseRange().Kvs[0].Value)
		fmt.Println("继续抢")
		time.Sleep(time.Second * 1)
	}
	return nil
}
