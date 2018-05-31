package amEtcd

import (
	"github.com/coreos/etcd/client"
	"os"
	"github.com/andy-zhangtao/AntMan/env"
	"time"
	"github.com/andy-zhangtao/AntMan/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"context"
)

var keysAPI client.KeysAPI

const (
	ModuleName = "AntMan-Etcd-Agent"
)

func init() {
	cfg := client.Config{
		Endpoints:               []string{os.Getenv(env.ANT_ENV_ETCD_ENDPOINT)},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5*time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		panic(log.Z.Error(fmt.Sprintf("ETCD Init Error [%s]", err.Error())))
	}

	keysAPI = client.NewKeysAPI(c)

	v, err := c.GetVersion(context.Background())
	if err != nil {
		panic(log.Z.Error(fmt.Sprintf("ETCD Connect Error [%s]", err.Error())))
	}
	logrus.WithFields(log.Z.Fields(logrus.Fields{"Etcd Init Success Server Version": v.Server, "Cluster Version": v.Cluster})).Info(ModuleName)
}
