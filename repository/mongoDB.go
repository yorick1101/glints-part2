package repository

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBOP struct {
	db *mongo.Database
}

func (op *MongoDBOP) GetCollection(name string) CollectionOP {
	return &MongoCollectionOP{
		collection: op.db.Collection(name),
	}
}

type MongoCollectionOP struct {
	collection *mongo.Collection
}

func (op *MongoCollectionOP) InsertMany(docs []interface{}) []interface{} {
	response, err := op.collection.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Panic(err)
	}
	return response.InsertedIDs
}

func (op *MongoCollectionOP) InsertOne(doc interface{}) interface{} {
	response, err := op.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Panic(err)
	}
	return response.InsertedID
}

func (op *MongoCollectionOP) UpsertOne(filter interface{}, update interface{}) interface{} {

	opts := options.Update().SetUpsert(true)

	res, err := op.collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Panic(err)
	}
	return res.UpsertedID
}

func (op *MongoCollectionOP) FindIdByAggregate(pipeline mongo.Pipeline) []string {
	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	cursor, err := op.collection.Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		log.Panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	var ids []string
	for _, result := range results {
		ids = append(ids, fmt.Sprintf("%v", result["_id"]))
	}
	return ids
}

func (op *MongoCollectionOP) Find(filter interface{}, results interface{}) {
	opts := options.Find()
	cursor, err := op.collection.Find(context.TODO(), filter, opts)
	defer cursor.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), results); err != nil {
		log.Fatal(err)
	}

}

func (op *MongoCollectionOP) FindWithOptions(filter interface{}, results interface{}, options *options.FindOptions) {

	cursor, err := op.collection.Find(context.TODO(), filter, options)
	defer cursor.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), results); err != nil {
		log.Fatal(err)
	}

}

func newMongoDBOP(username string, password string, host string, port int, databaseName string) *MongoDBOP {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	connectionURI := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI).SetAuth(credential))
	if err != nil {
		log.Panic("failed to connect to database", err)
	}
	log.Info("db connected")
	database := client.Database(databaseName)
	log.Info("db switch database", database.Name())

	return &MongoDBOP{
		db: database,
	}
}
