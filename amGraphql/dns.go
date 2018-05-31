package amGraphql

import (
	"github.com/graphql-go/graphql"
	"github.com/andy-zhangtao/AntMan/amEtcd"
	"github.com/andy-zhangtao/gogather/strings"
	"github.com/sirupsen/logrus"
	"github.com/andy-zhangtao/AntMan/log"
	"fmt"
	"encoding/json"
	"github.com/andy-zhangtao/AntMan/model"
)

var DnsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "dns_object",
	Fields: graphql.Fields{
		"domain": &graphql.Field{
			Type:        graphql.String,
			Description: "The DNS record domain",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if d, ok := p.Source.(model.Dns); ok {

					return strings.ReverseWithSeg(d.Domain, "/", "."), nil
				}
				return nil, nil
			},
		},
		"address": &graphql.Field{
			Type:        graphql.String,
			Description: "The domain ip address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if d, ok := p.Source.(model.Dns); ok {

					type h struct {
						Host string `json:"host"`
					}

					var _h h
					err := json.Unmarshal([]byte(d.Address), &_h)
					if err != nil {
						return nil, log.Z.Error(fmt.Sprintf("Unmarshal Addess Error [%s]", err.Error()))
					}

					return _h.Host, nil
				}
				return nil, nil
			},
		},
	},
})

var DnsQuery = &graphql.Field{
	Type:        graphql.NewList(DnsType),
	Description: "Query All/Specify Dns Record",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return amEtcd.GetAllKeys()
	},
}

var DnsNew = &graphql.Field{
	Type:        DnsType,
	Description: "Add a new Dns record",
	Args: graphql.FieldConfigArgument{
		"domain": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Dns access domain",
		},
		"address": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Dns access address",
		},
		"env": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Dns's environment",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		domain, _ := p.Args["domain"].(string)
		address, _ := p.Args["address"].(string)
		env, _ := p.Args["env"].(string)

		reverDomain := strings.ReverseWithSeg(domain, ".", "/")

		reverDomain = fmt.Sprintf("/%s/%s", env, reverDomain)
		type host struct {
			Host string `json:"host"`
		}

		data, err := json.Marshal(&host{
			Host: address,
		})

		if err != nil {
			logrus.WithFields(logrus.Fields{"Error": log.Z.Error(fmt.Sprintf("Marshal domain Error [%s]", err.Error()))}).Error(ModuleName)
			return nil, err
		}

		logrus.WithFields(log.Z.Fields(logrus.Fields{"domain": reverDomain, "address": string(data)})).Info(ModuleName)

		if err := amEtcd.AddNewEntry(reverDomain, string(data)); err != nil {
			logrus.WithFields(logrus.Fields{"Error": log.Z.Error(fmt.Sprintf("Add new domain Error [%s]", err.Error()))}).Error(ModuleName)
			return nil, err
		}

		return nil, nil
	},
}

var DnsDelete = &graphql.Field{
	Type:        DnsType,
	Description: "Delete a existing dns record",
	Args: graphql.FieldConfigArgument{
		"domain": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Dns access domain",
		},
		"env": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Dns's environment",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		domain, _ := p.Args["domain"].(string)
		env, _ := p.Args["env"].(string)

		domain = fmt.Sprintf("%s.%s.", domain, env)

		domain = strings.ReverseWithSeg(domain, ".", "/")
		logrus.WithFields(log.Z.Fields(logrus.Fields{"Delete Domain": domain})).Info(ModuleName)

		return nil, amEtcd.DeleteKeys(domain)
	},
}
