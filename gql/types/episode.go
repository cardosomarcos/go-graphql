package types

import "github.com/graphql-go/graphql"

var Episode = graphql.NewObject(graphql.ObjectConfig{
	Name: "episode",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "episode id",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "episode title",
		},
	}})
