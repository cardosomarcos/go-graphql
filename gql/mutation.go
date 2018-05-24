package gql

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Mutation",
		Fields: graphql.Fields{
			"createAuthor": &graphql.Field{
				Type:        types.Author,
				Description: "create author",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.CreateAuthor,
			},
			"createEpisode": &graphql.Field{
				Type:        types.Episode,
				Description: "create episode",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.CreateEpisode,
			},
		}})
