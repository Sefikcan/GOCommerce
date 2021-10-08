package Connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ordering/common/constants"
)

var MongoCollection *mongo.Collection

func Connect(){
	clientOptions := options.Client().ApplyURI(constants.ORDERINGCONNSTRING)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	MongoCollection = client.Database(constants.ORDERINGDB).Collection(constants.ORDERINGCL)
}
