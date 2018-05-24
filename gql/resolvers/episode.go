package resolvers

import (
	"go-graphql/gql/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

type Episode struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Title string        `json:"title"`
}

func GetEpisode(params graphql.ResolveParams) (interface{}, error) {
	episodeid := params.Args["id"].(string)
	var res Episode
	_ = db.Mongo.DB("demo").C("episode").Find(bson.M{"_id": bson.ObjectIdHex(episodeid)}).One(&res)
	return res, nil
}

func CreateEpisode(params graphql.ResolveParams) (interface{}, error) {
	title := params.Args["title"].(string)
	id := bson.NewObjectId()

	episode := Episode{
		Id:    id,
		Title: title,
	}
	_ = db.Mongo.DB("demo").C("episode").Insert(&episode)
	return &episode, nil
}
