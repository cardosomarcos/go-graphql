package fields

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

var EpisodeQuery = &graphql.Field{
	Type:        types.Author,
	Description: "get author",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.GetAuthor,
}

var CreateEpisode = &graphql.Field{
	Type:        types.Episode,
	Description: "create episode",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.CreateEpisode,
}
