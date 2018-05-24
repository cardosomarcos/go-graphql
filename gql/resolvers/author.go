package resolvers

import (
	"go-graphql/gql/db"
	"time"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

type Author struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name"`
	Age      int           `json:"age"`
	UpdateAt time.Time     `json:"updateAt"`
}

func GetAuthor(params graphql.ResolveParams) (interface{}, error) {
	authorid := params.Args["id"].(string)
	var res Author
	_ = db.Mongo.DB("demo").C("author").Find(bson.M{"_id": bson.ObjectIdHex(authorid)}).One(&res)
	return res, nil
}

func CreateAuthor(params graphql.ResolveParams) (interface{}, error) {
	id := bson.NewObjectId()
	name := params.Args["name"].(string)
	age := params.Args["age"].(int)
	author := Author{
		Id:       id,
		Name:     name,
		Age:      age,
		UpdateAt: time.Now(),
	}
	_ = db.Mongo.DB("demo").C("author").Insert(&author)
	return author, nil
}
