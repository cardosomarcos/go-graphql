package resolvers

import (
	"go-graphql/gql/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

type Episode struct {
	Id    bson.ObjectId `json:"id"`
	Title string        `json:"title"`
}

func GetEpisode(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(bson.ObjectId)
	var res Episode
	_ = db.Mongo.DB("demo").C("episode").Find(bson.M{"id": id}).One(&res)
	return res, nil
}
