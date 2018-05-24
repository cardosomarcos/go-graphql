package gql

import (
	"go-graphql/gql/fields"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "rootQuery",
		Description: "rootQuery",
		Fields: graphql.Fields{
			"episode": fields.EpisodeQuery,
			"author":  fields.AuthorQuery,
		},
	},
)
