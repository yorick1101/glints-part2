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
	FindIdByAggregate(pipeline mongo.Pipeline) ([]string, error)
	Find(filter interface{}, results interface{}) error
	FindOne(filter interface{}, results interface{}) error
	FindWithOptions(filter interface{}, results interface{}, options *options.FindOptions) error
	Replace(id string, update interface{}) (int64, error)
}

func newDBOP(username string, password string, host string, port int, databaseName string) DBOP {
	return newMongoDBOP(username, password, host, port, databaseName)
}
