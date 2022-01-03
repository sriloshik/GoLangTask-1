package db

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

func GetCollection(server, bucketName, user, password, scopeName, collectionName string) (*gocb.Collection, error) {

	cluster, err := gocb.Connect(
		server,
		gocb.ClusterOptions{
			Username: user,
			Password: password,
		})
	if err != nil {
		return nil, err
	}

	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}

	scope := bucket.Scope(scopeName)
	collection := scope.Collection(collectionName)

	return collection, nil
}
