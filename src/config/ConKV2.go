package config

import (
	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB : GetMongoDB
func GetMongoDB() (*mgo.Database, error) {
	host := "mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"
	dbName := "KV2MongoDB"
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)
	return db, nil
}
