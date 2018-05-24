package main

import (
	"go-graphql/gql"
	"go-graphql/gql/db"
	"log"
	"net/http"
	"runtime"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
	mgo "gopkg.in/mgo.v2"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	db.Mongo, _ = mgo.Dial("mongodb://127.0.0.1:27017")
	db.Mongo.SetMode(mgo.Monotonic, true)
	db.Mongo.SetPoolLimit(600)
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
