package types

import (
	"github.com/graphql-go/graphql"
)

var Author = graphql.NewObject(graphql.ObjectConfig{
	Name: "author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "author id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "author name",
		},
		"age": &graphql.Field{
			Type:        graphql.Int,
			Description: "author age",
		},
	}})
