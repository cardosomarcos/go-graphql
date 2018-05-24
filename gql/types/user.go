package types

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "user name",
		},
	}})
