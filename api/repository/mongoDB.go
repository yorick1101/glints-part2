package repository

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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

func (op *MongoCollectionOP) FindIdByAggregate(pipeline mongo.Pipeline) ([]string, error) {
	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	var ids []string
	cursor, err := op.collection.Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		return ids, err
	}

	var results []IdContainer
	if err = cursor.All(context.TODO(), &results); err != nil {
		return ids, err
	}

	for _, result := range results {
		ids = append(ids, result.Id.Hex())
	}
	return ids, nil
}

//results must be pointer to array
func (op *MongoCollectionOP) Find(filter interface{}, results interface{}) error {
	opts := options.Find()
	cursor, err := op.collection.Find(context.TODO(), filter, opts)
	log.Info("filter:", filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), results); err != nil {
		return err
	}
	return nil
}

func (op *MongoCollectionOP) FindOne(filter interface{}, result interface{}) error {
	opts := options.FindOne()

	singleResult := op.collection.FindOne(context.TODO(), filter, opts)
	err := singleResult.Err()
	if err != nil {
		return err
	}
	err = singleResult.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (op *MongoCollectionOP) FindWithOptions(filter interface{}, results interface{}, options *options.FindOptions) error {

	cursor, err := op.collection.Find(context.TODO(), filter, options)
	defer cursor.Close(context.TODO())
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), results); err != nil {
		return err
	}
	return nil
}

func (op *MongoCollectionOP) Replace(filter interface{}, update interface{}) (int64, error) {
	opts := options.Replace().SetUpsert(false)
	res, err := op.collection.ReplaceOne(context.TODO(), filter, update, opts)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (op *MongoCollectionOP) Delete(filter interface{}) (int64, error) {
	opts := options.Delete()
	res, err := op.collection.DeleteMany(context.TODO(), filter, opts)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
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
	log.Info("db switch database ", database.Name())

	return &MongoDBOP{
		db: database,
	}
}
