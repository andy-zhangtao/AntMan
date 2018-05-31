package amNsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/andy-zhangtao/AntMan/env"
	"fmt"
	"encoding/json"
	"github.com/andy-zhangtao/AntMan/log"
	"github.com/andy-zhangtao/AntMan/model"
)

type DataAgent struct{}

func (this *DataAgent) HandleMessage(m *nsq.Message) error {
	logrus.WithFields(log.Z.Fields(logrus.Fields{"HandleMessage": string(m.Body)})).Info(ModelName)
	m.DisableAutoResponse()
	workChan <- m
	return nil
}

func HandlerServiceNsq() (err error) {
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = 1000

	r, err := nsq.NewConsumer(os.Getenv(env.ENV_NSQ_SVC_TOPIC), ModelName, cfg)
	if err != nil {
		return log.Z.Error(fmt.Sprintf("Connect Nsq Build Topic Error [%s]", err.Error()))
	}

	go func() {
		logrus.WithFields(log.Z.Fields(logrus.Fields{"WorkChan Listen Status": os.Getenv(env.ENV_NSQ_SVC_TOPIC)})).Info(ModelName)
		for m := range workChan {
			logrus.WithFields(log.Z.Fields(logrus.Fields{"Recevie Build Request": string(m.Body)})).Info(ModelName)
			req := model.MsgEvent{}

			err = json.Unmarshal(m.Body, &req)
			if err != nil {
				logrus.WithFields(logrus.Fields{"Unmarshal Request Error": err}).Error(ModelName)
				continue
			}

			go handlerMsgEvent(req)
			m.Finish()
		}
	}()
	r.AddConcurrentHandlers(&DataAgent{}, 20)
	err = r.ConnectToNSQD(os.Getenv(env.ENV_NSQ_ENDPOINT))
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	return nil
}

func handlerMsgEvent(req model.MsgEvent) {
	switch req.Kind {
	case env.SERVICE_CHANGE:
		//service := req.Content.(model.Change)
		return
	}
}
