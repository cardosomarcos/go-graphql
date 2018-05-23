package types

import "github.com/graphql-go/graphql"

var Child = graphql.NewObject(graphql.ObjectConfig{
	Name: "child",
	Fields: graphql.Fields{
		"entity_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "qipu_id",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "episode title",
		},
	}})

var Episode = graphql.NewObject(graphql.ObjectConfig{
	Name: "episode",
	Fields: graphql.Fields{
		"entity_id": &graphql.Field{
			Type:        graphql.Int,
			Description: "qipu_id",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "episode title",
		},
	}})
