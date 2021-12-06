package gimgraphql

import "github.com/onichandame/gim"

var GraphqlModule = gim.Module{
	Name:      "GraphqlModule",
	Providers: []interface{}{newGraphqlService, newParser},
	Exports:   []interface{}{newGraphqlService, newParser},
}
