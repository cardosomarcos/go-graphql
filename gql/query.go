package gql

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "rootQuery",
		Description: "rootQuery",
		Fields: graphql.Fields{
			"episode": &graphql.Field{
				Type:        types.Episode,
				Description: "get episode",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetEpisode,
			},
			"author": &graphql.Field{
				Type:        types.Author,
				Description: "get author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.GetAuthor,
			},
		},
	},
)
