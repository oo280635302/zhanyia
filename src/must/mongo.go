package must

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	URL = "mongodb://localhost:27017"
)

func GetMongoDB() (*mongo.Client, func()) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URL))
	if err != nil {
		fmt.Println("CLIENT ERROR:", err)
		return nil, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("PING ERROR:", err)
		return nil, nil
	}
	return client, func() {
		client.Disconnect(ctx)
		cancel()
	}
}

func Fmttt() (int, func()) {
	a := 11230
	return 1, func() {
		fmt.Println(a)
	}
}
