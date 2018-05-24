package resolvers

import (
	"fmt"
	"go-graphql/gql/db"
	"time"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

type Author struct {
	Id         int       `json:"id" bson:"_id"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	LastChange time.Time `json:"lastChange"`
	UpdateAt   time.Time `json:"updateAt"`
}

func GetAuthor(params graphql.ResolveParams) (interface{}, error) {
	authorid := params.Args["id"].(int)
	var res Author
	_ = db.Mongo.DB("demo").C("author").Find(bson.M{"id": authorid}).One(&res)
	return res, nil
}

func CreateAuthor(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	fmt.Println("chegou aqui")
	var res Author
	res.Name = name
	res.Id = 2
	//_ = db.Mongo.DB("demo").C("author").Find(bson.M{"id": authorid}).One(&res)
	return res, nil
}
