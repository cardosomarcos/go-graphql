package resolvers

import (
	"go-graphql/gql/db"
	"time"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name"`
	Age        int           `json:"age"`
	LastChange time.Time     `json:"lastChange"`
	UpdateAt   time.Time     `json:"updateAt"`
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	userid := params.Args["userid"].(int)
	var res User
	_ = db.Mongo.DB("demo").C("user").Find(bson.M{"userid": userid}).One(&res)
	return res, nil
}
