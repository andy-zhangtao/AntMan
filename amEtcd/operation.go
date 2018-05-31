package amEtcd

import (
	"context"
	"github.com/andy-zhangtao/AntMan/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"github.com/andy-zhangtao/AntMan/model"
)

func AddNewEntry(key, value string) (error) {
	resp, err := keysAPI.Set(context.Background(), key, value, nil)
	if err != nil {
		return log.Z.Error(fmt.Sprintf("Insert New Key Error [%s]", err.Error()))
	}

	logrus.WithFields(log.Z.Fields(logrus.Fields{"Insert New Key Resp ": resp})).Info(ModuleName)

	return nil
}

func GetAllKeys() (dns []model.Dns, err error) {
	logrus.WithFields(log.Z.Fields(logrus.Fields{"Query Etcd Directory": "/"})).Info(ModuleName)

	return getDirectValues("/")
}

func getDirectValues(dir string) (dns []model.Dns, err error) {
	logrus.WithFields(log.Z.Fields(logrus.Fields{"Query Etcd Directory": dir})).Info(ModuleName)
	resp, err := keysAPI.Get(context.Background(), dir, nil)
	if err != nil {
		err = log.Z.Error(fmt.Sprintf("Query Etcd Directory [%s] Error [%s]", dir, err.Error()))
		return
	}

	sort.Sort(resp.Node.Nodes)
	for _, n := range resp.Node.Nodes {
		if n.Dir {
			_dns, err := getDirectValues(n.Key)
			if err != nil {
				return nil, err
			}
			dns = append(dns, _dns...)
		} else {
			dns = append(dns, model.Dns{
				Domain:  n.Key,
				Address: n.Value,
			})
		}
	}

	return
}

func DeleteKeys(key string) error {
	resp, err := keysAPI.Delete(context.Background(), key, nil)
	if err != nil {
		return log.Z.Error(fmt.Sprintf("Delete Key [%s] Error [%s]", key, err.Error()))
	}

	logrus.WithFields(log.Z.Fields(logrus.Fields{"Delete Key Resp": resp})).Info(ModuleName)
	return nil
}
