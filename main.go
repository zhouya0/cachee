package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	// 为了保证 HTTPS 链接可信，需要预先加载目标证书签发机构的 CA 根证书
	etcdCA, err := ioutil.ReadFile("/etc/kubernetes/pki/etcd/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair("/etc/kubernetes/pki/etcd/server.crt", "/etc/kubernetes/pki/etcd/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个空的 CA Pool
	// 因为后续只会链接 Etcd 的 api 端点，所以此处选择使用空的 CA Pool，然后只加入 Etcd CA 既可
	// 如果期望链接其他 TLS 端点，那么最好使用 x509.SystemCertPool() 方法先 copy 一份系统根 CA
	// 然后再向这个 Pool 中添加自定义 CA
	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(etcdCA)

	// 创建 api v3 的 client
	cli, err := clientv3.New(clientv3.Config{
		// etcd https api 端点
		Endpoints:   []string{"https://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		// 自定义 CA 及 Client Cert 配置
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = cli.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	putResp, err := cli.Put(ctx, "sample_key", "sample_value")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(putResp)
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	delResp, err := cli.Delete(ctx, "sample_key")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(delResp)
	}
	cancel()
}
