package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"wxcloudrun-golang/db"
)

const tableName = "todo"

func (imp TodoInterfaceImp) SaveTodos(openid string, todos interface{}) error {
	collection := db.GetMongoDb().Collection(tableName)
	m := map[string]interface{}{
		"_id":   openid,
		"todos": todos,
	}
	filter := bson.D{{"_id", openid}}
	collection.DeleteOne(context.TODO(), filter)
	collection.InsertOne(context.TODO(), m)
	return nil
}

func (imp TodoInterfaceImp) GetTodos(openid string) (todos interface{}, err error) {
	collection := db.GetMongoDb().Collection(tableName)
	var op interface{} = openid
	filter := bson.D{{"_id", op}}
	var result interface{}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}
