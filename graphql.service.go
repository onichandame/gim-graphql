package graphql

import (
	"github.com/graphql-go/graphql"
	goutils "github.com/onichandame/go-utils"
)

type GraphqlService struct {
	queries       graphql.Fields
	mutations     graphql.Fields
	subscriptions graphql.Fields
}

func newGraphqlService() *GraphqlService {
	var svc GraphqlService
	svc.queries = make(graphql.Fields)
	svc.mutations = make(graphql.Fields)
	svc.subscriptions = make(graphql.Fields)
	return &svc
}

func (svc *GraphqlService) AddQuery(name string, query *graphql.Field) {
	svc.queries[name] = query
}
func (svc *GraphqlService) AddMutation(name string, mutation *graphql.Field) {
	svc.mutations[name] = mutation
}
func (svc *GraphqlService) AddSubscription(name string, subscription *graphql.Field) {
	svc.subscriptions[name] = subscription
}

func (svc *GraphqlService) BuildSchema() *graphql.Schema {
	var schemaConf graphql.SchemaConfig
	if len(svc.queries) > 0 {
		schemaConf.Query = graphql.NewObject(graphql.ObjectConfig{
			Name:   `Query`,
			Fields: svc.queries,
		})
	}
	if len(svc.mutations) > 0 {
		schemaConf.Mutation = graphql.NewObject(graphql.ObjectConfig{
			Name:   `Mutation`,
			Fields: svc.mutations,
		})
	}
	if len(svc.subscriptions) > 0 {
		schemaConf.Subscription = graphql.NewObject(graphql.ObjectConfig{
			Name:   `Subscription`,
			Fields: svc.subscriptions,
		})
	}
	schema, err := graphql.NewSchema(schemaConf)
	goutils.Assert(err)
	return &schema
}
