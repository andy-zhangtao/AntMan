package amNsq

import (
	"github.com/nsqio/go-nsq"
		)

const ModelName = "AntMan-Nsq-Agent"

var workChan chan *nsq.Message

func init() {
	workChan = make(chan *nsq.Message)
}
