# Golang, GraphQL e MongoDB

Testing GraphQL API implemented in Go and MongoDB

## Todo 

- List all authors
- List all episodes
- ~~Create mutations~~
- Authorization
- Migrate to serverless 


## Commands
This application expose a single endpoints '/graphql' wich accepts both queries and mutations


- Get all fields
```
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n\tauthor(id: \"5b071dc428c1e324878952b7\"){\n\t\tid, name, age\n\t}\n\t\n\tepisode(id: \"5b0723d828c1e33e9e3d8baf\"){\n\t\tid, title\n\t}\n\t\n}"}'
```

- Get episode
``` 
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n\n\tepisode(id: \"5b0723d828c1e33e9e3d8baf\"){\n\t\tid, title\n\t}\n\t\n}"}'
```

- Get author
``` 
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n\tauthor(id: \"5b071dc428c1e324878952b7\"){\n\t\tid, name, age\n\t}\n\t\n\n\t\n}"}'
```
- Insert author 
```
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"mutation {\n\tcreateAuthor (name: \"Cardoso Marcos\", age: 22){\n\t\tid, name, age\n\t}\n\t\n}"}'
```

- Insert episode
```
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"mutation {\n\tcreateEpisode (title: \"The life of Pable\"){\n\t\tid, title\n\t}\n\t\n}"}'
```