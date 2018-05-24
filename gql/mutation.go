package gql

import (
	"go-graphql/gql/fields"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Mutation",
		Fields: graphql.Fields{
			"createAuthor":  fields.CreateAuthor,
			"createEpisode": fields.CreateEpisode,
		}})
