package repository

import "go.mongodb.org/mongo-driver/mongo"

type DBOP interface {
	GetCollection(name string) CollectionOP
}

type CollectionOP interface {
	InsertMany(docs []interface{}) []interface{}
	InsertOne(doc interface{}) interface{}
	UpsertOne(filter interface{}, update interface{}) interface{}
	FindIdByAggregate(pipeline mongo.Pipeline) []string
	Find(filter interface{}) []interface{}
}

func newDBOP(username string, password string, host string, port int, databaseName string) DBOP {
	return newMongoDBOP(username, password, host, port, databaseName)
}
