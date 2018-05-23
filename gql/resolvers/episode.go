package resolvers

import (
	"go-graphql/gql/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

type Episode struct {
	Entity_id int    `json:"entity_id"`
	Title     string `json:"title"`
}

func GetEpisode(params graphql.ResolveParams) (interface{}, error) {
	entity_id := params.Args["entity_id"].(int)
	var res Episode
	_ = db.Mongo.DB("demo").C("episode").Find(bson.M{"entity_id": entity_id}).One(&res)
	return res, nil
}
