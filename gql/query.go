package gql

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

func GetEpisode() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "episode",
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
			},
		},
	)
}

func GetUser() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
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
		},
	})
}
