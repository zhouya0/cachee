package client

import (
	"cachee/config"
	"crypto/tls"
	"crypto/x509"
	"go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"log"
	"time"
)

// GetETCDClient will return the etcd client with https certificates.
func GetETCDClient() *clientv3.Client {
	var c config.Config
	config := c.GetConfig()

	etcdCA, err := ioutil.ReadFile(config.CAPath)
	if err != nil {
		log.Fatal(err)
	}

	etcdClientCert, err := tls.LoadX509KeyPair(config.CertPath, config.KeyPath)
	if err != nil {
		log.Fatal(err)
	}

	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(etcdCA)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.Server},
		DialTimeout: 5 * time.Second,
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return cli
}
