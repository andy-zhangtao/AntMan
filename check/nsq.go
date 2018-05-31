package check

import (
	"os"
	"github.com/andy-zhangtao/AntMan/env"
	"fmt"
	"github.com/andy-zhangtao/AntMan/log"
)

func checkNSQ() error {
	if os.Getenv(env.ENV_NSQ_ENDPOINT) == "" {
		return log.Z.Error(fmt.Sprintf("[%s] Empty!", env.ENV_NSQ_ENDPOINT))
	}

	if os.Getenv(env.ENV_NSQ_SVC_TOPIC) == "" {
		return log.Z.Error(fmt.Sprintf("[%s] Empty!", env.ENV_NSQ_SVC_TOPIC))
	}

	return nil
}


