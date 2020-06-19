# Rest API MongoDB

docker-compose up

create database people over localhost:8081

## Dependencies

go get github.com/gorilla/mux
go get github.com/mongodb/mongo-go-driver/mongo

## Insomnia

Get People

- http://localhost:12345/people

Post person

```
{
	"firstName": "Mike",
	"lastName": "Ze"
}
```

Get Person
- http://localhost:12345/person/\<id-from-post>

