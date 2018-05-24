package gql

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

func GetRootQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "rootQuery",
			Description: "rootQuery",
			Fields: graphql.Fields{
				"episode": &graphql.Field{
					Type:        types.Episode,
					Description: "get episode",
					Args: graphql.FieldConfigArgument{
						"entity_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: resolvers.GetEpisode,
				},
				"user": &graphql.Field{
					Type:        types.User,
					Description: "get user",
					Args: graphql.FieldConfigArgument{
						"userid": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: resolvers.GetUser,
				},
				"users" 
			},
		},
	)
}
