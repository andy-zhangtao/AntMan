package check

import (
	"os"
	"github.com/andy-zhangtao/AntMan/env"

	"fmt"
	"github.com/andy-zhangtao/AntMan/log"
)

func checkEtcd() error {
	if os.Getenv(env.ANT_ENV_ETCD_ENDPOINT) == "" {
		return log.Z.Error(fmt.Sprintf("[%s] Empty!", env.ANT_ENV_ETCD_ENDPOINT))
	}

	//if os.Getenv(env.ANT_ENV_ETCD_CLUSTER_CHAIN) == "" {
	//	return log.Z.Error(fmt.Sprintf("[%s] Empty!", env.ANT_ENV_ETCD_CLUSTER_CHAIN))
	//}

	return nil
}
