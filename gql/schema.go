package gql

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    Query,
	Mutation: Mutation,
	
})
