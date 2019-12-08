package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBOP interface {
	GetCollection(name string) CollectionOP
}

type CollectionOP interface {
	InsertMany(docs []interface{}) []interface{}
	InsertOne(doc interface{}) interface{}
	UpsertOne(filter interface{}, update interface{}) interface{}
	FindIdByAggregate(pipeline mongo.Pipeline) []string
	Find(filter interface{}, results interface{})
	FindWithOptions(filter interface{}, results interface{}, options *options.FindOptions)
}

func newDBOP(username string, password string, host string, port int, databaseName string) DBOP {
	return newMongoDBOP(username, password, host, port, databaseName)
}
