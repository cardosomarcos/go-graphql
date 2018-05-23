package gql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func GraphQlHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}

func HandlerGQL() {
}
