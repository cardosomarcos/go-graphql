## Golang + graphQL

### Test graphql 

Using insomnia: 
http://localhost:3000/graphql

curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"fragment WithPost on Post{\n\tid\n  title\n  body\n}\n\nfragment Comment on Comment {\n\tid\n  email\n  name\n}\n\nfragment WithComment on Post {\n\t...Comment\n}\n\n{\n post(id: 5) {\n\t\t\t...WithPost\t\n\t\t\t...WithComment\n\t}\n}"}'

