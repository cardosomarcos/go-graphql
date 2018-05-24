## Golang + graphQL

### Test graphql 

Using insomnia: 
http://localhost:3000/graphql


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