package fields

import (
	"go-graphql/gql/resolvers"
	"go-graphql/gql/types"

	"github.com/graphql-go/graphql"
)

var AuthorQuery = &graphql.Field{
	Type:        types.Episode,
	Description: "get episode",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.GetAuthor,
}

var CreateAuthor = &graphql.Field{
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
}
