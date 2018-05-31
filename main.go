package main

import (
	_ "github.com/andy-zhangtao/AntMan/check"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"github.com/andy-zhangtao/AntMan/log"
	"github.com/andy-zhangtao/AntMan/amGraphql"
	"github.com/rs/cors"
)

const ModuleName = "AntMan-Main-Agent"

func main() {
	router := mux.NewRouter()
	router.Path("/api").HandlerFunc(handleDevExGraphQL)
	handler := cors.AllowAll().Handler(router)
	logrus.Fatal(http.ListenAndServe(":8000", handler))
}

var schemaDevex, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootDnsQuery,
	Mutation: rootMutation,
})

var rootDnsQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"dns": amGraphql.DnsQuery,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addDns":    amGraphql.DnsNew,
		"deleteDns": amGraphql.DnsDelete,
	},
})

func handleDevExGraphQL(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var g map[string]interface{}
	if r.Method == http.MethodGet {
		g = make(map[string]interface{})
		g["query"] = r.URL.Query().Get("query")
		logrus.WithFields(log.Z.Fields(logrus.Fields{"query": g["query"]})).Info(ModuleName)
		result := executeDevExQuery(g, schemaDevex)
		json.NewEncoder(w).Encode(result)
	}

	if r.Method == http.MethodPost {
		data, _ := ioutil.ReadAll(r.Body)

		err := json.Unmarshal(data, &g)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
		result := executeDevExQuery(g, schemaDevex)
		json.NewEncoder(w).Encode(result)
	}
}

func executeDevExQuery(query map[string]interface{}, schema graphql.Schema) *graphql.Result {

	params := graphql.Params{
		Schema:        schema,
		RequestString: query["query"].(string),
	}

	if query["variables"] != nil {
		params.VariableValues = query["variables"].(map[string]interface{})
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		logrus.WithFields(log.Z.Fields(logrus.Fields{"wrong result, unexpected errors": result.Errors})).Error(ModuleName)

	}
	return result
}
