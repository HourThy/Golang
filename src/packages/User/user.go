package user

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User : User struct
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username, omitempty"`
	Password string             `bson:"password, omitempty"`
}

// InsertUser : Create a new user
func InsertUser(username, password string) string {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		println(err)
	}
	fmt.Println("Connected to MongoDb!")

	collection := client.Database("KV2MongoDB").Collection("dbLogin")
	filter := bson.D{{"username", username}, {"password", password}}
	insertResult, err := collection.InsertOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	println("Inserted to database.")
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecting to MongoDb closed.")
	fmt.Println("Inserted many document: ", insertResult.InsertedID)
	return "User inserted successful!"
}

// DeleteUser : Delete
func DeleteUser(id string) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("User")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

//UpdateRecord : Update
func UpdateRecord(id, username, password string) {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("dbLogin").Collection("User")
	println("Collection type:", reflect.TypeOf(collection))
	docID := id
	objID, err := primitive.ObjectIDFromHex(docID)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", objID)
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.D{{"username", username}, {"password", password}}}
	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}
}
//machno, matetype, matecode, matename, intime, indat, usrno, state string
//{"machno", machno}, {"matetype", matetype}, {"matecode", matecode}, {"matename", matename}, {"intime", intime}, {"indat", indat}, {"usrno", usrno}, {"state", state}