# Golang, GraphQL e MongoDB

Testing GraphQL API implemented in Go and MongoDB

## Todo 

- List all authors
- List all episodes
- Create mutations
- Authorization


## Commands
This application expose a single endpoints '/graphql' wich accepts both queries and mutations


- Get all fields
```
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n  episode(id: 1) {\n\t\ttitle\n\t}\n\tauthor(id: 1){\n\t\tname\n\t}\n}"}'
```

- Get episode field
``` 
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n  episode(id: 1) {\n\t\ttitle\n\t}\n}"}'
```

- Get author field
``` 
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"{\n \tauthor(id: 1){\n\t\tname\n\t}\n}"}'
```
- Insert author 
```
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"mutation {\n\tcreateAuthor (name: \"Cardoso Marcos\", age: 22){\n\t\tid, name, age\n\t}\n\t\n}"}'
```