package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"logbyte/src/types"
)

var logbyte *mongo.Collection

func Init(config *types.ConfigDb) (db *mongo.Database, err error) {
	var client *mongo.Client
	client, err = mongo.NewClient(options.Client().ApplyURI(config.URL))
	if err != nil {
		return
	}

	err = client.Connect(context.Background())
	if err != nil {
		return
	}

	db = client.Database(config.Database)
	logbyte = db.Collection(config.Collection)
	return
}

func Get(id string) (data *types.LogEntry, err error) {
	res := logbyte.FindOne(context.Background(), _id(id))
	if res.Err() != nil {
		return
	}

	err = res.Decode(&data)
	return
}

func Delete(id string) (data *types.LogEntry, err error) {
	data, err = Get(id)
	if err != nil {
		return
	}

	_, err = logbyte.DeleteOne(context.Background(), _id(id))
	return
}

func Insert(notif *types.LogEntry) (err error) {
	_, err = logbyte.InsertOne(context.Background(), notif)
	return
}

func Patch(id string, json []byte) (data *types.LogEntry, err error) {
	var doc interface{}
	err = bson.UnmarshalJSON(json, &doc)

	_, err = logbyte.UpdateOne(context.Background(), _id(id), bson.M{"$set": doc})
	if err != nil {
		return
	}

	return Get(id)
}

func Paginate(limit int64, offset int64) ([]*types.LogEntry, error) {
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	opts.SetSkip(limit * offset)
	opts.SetLimit(limit)

	find, err := logbyte.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer find.Close(context.Background())

	data := make([]*types.LogEntry, 0)
	for find.Next(context.Background()) {
		var result *types.LogEntry
		err = find.Decode(&result)
		if err != nil {
			return nil, err
		}

		data = append(data, result)
	}

	err = find.Err()
	return data, err
}

func Count() (count int64, err error) {
	return logbyte.CountDocuments(context.Background(), bson.M{})
}

func _id(id string) bson.M {
	return bson.M{"_id": id}
}
