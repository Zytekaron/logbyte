package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"notifs/src/types"
)

var notifs *mongo.Collection

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
	notifs = db.Collection(config.Collection)
	return
}

func Get(id string) (data *types.Notification, err error) {
	res := notifs.FindOne(context.Background(), _id(id))
	if res.Err() != nil {
		return
	}

	err = res.Decode(&data)
	return
}

func Delete(id string) (data *types.Notification, err error) {
	data, err = Get(id)
	if err != nil {
		return
	}

	_, err = notifs.DeleteOne(context.Background(), _id(id))
	return
}

func Insert(notif *types.Notification) (err error) {
	_, err = notifs.InsertOne(context.Background(), notif)
	return
}

func Patch(id string, json []byte) (data *types.Notification, err error) {
	var doc interface{}
	err = bson.UnmarshalJSON(json, &doc)

	_, err = notifs.UpdateOne(context.Background(), _id(id), bson.M{"$set": doc})
	if err != nil {
		return
	}

	return Get(id)
}

func _id(id string) bson.M {
	return bson.M{"_id": id}
}
