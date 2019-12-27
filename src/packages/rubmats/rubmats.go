package rubmats

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Rubmats : Rubmats struct
type Rubmats struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Machno   string             `bson:"machno, omitempty"`
	Matetype string             `bson:"matetype, omitempty"`
	Matecode string             `bson:"matecode, omitempty"`
	Matename string             `bson:"matename, omitempty"`
	Intime   string             `bson:"intime, omitempty"`
	Indat    string             `bson:"indat, omitempty"`
	Usrno    string             `bson:"usrno, omitempty"`
	State    string             `bson:"state, omitempty"`
}

// InsertRubmats : Create a new us
func InsertRubmats(machno, matetype, matecode, matename, intime, indat, usrno, state string) string {
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

	collection := client.Database("KV2MongoDB").Collection("rubmats")
	filter := bson.D{{"machno", machno}, {"matetype", matetype}, {"matecode", matecode}, {"matename", matename}, {"intime", intime}, {"indat", indat}, {"usrno", usrno}, {"state", state}}
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
	return "Rubmats inserted successful!"
}

// DeleteRubmats : DeleteRubmats
func DeleteRubmats(id string) {
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
	collection := client.Database("KV2MongoDB").Collection("rubmats")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

//UpdateRubmats : Update
func UpdateRubmats(id, machno, matetype, matecode, matename, intime, indat, usrno, state string) {
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
	collection := client.Database("KV2MongoDB").Collection("rubmats")
	println("Collection type:", reflect.TypeOf(collection))
	docID := id
	objID, err := primitive.ObjectIDFromHex(docID)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", objID)
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.D{{"machno", machno}, {"matetype", matetype}, {"matecode", matecode}, {"matename", matename}, {"intime", intime}, {"indat", indat}, {"usrno", usrno}, {"state", state}}}
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

//FindSingleMachno : FindSingleMachno
func FindSingleMachno(Machno string) {
	// Set client opions
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")

	// Connect to MngoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println("Connecing to MongoDB failed!")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("rubmats")
	filter := bson.D{{"machno", Machno}}
	var results Rubmats
	err = collection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		println("Find achno error")
	}
	fmt.Printf("Foud a single document: %+v\n", results)
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connectio to MongoDB close.")
}

//FindAllMachno : FindAllMachno
func FindAllMachno() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("rubmats")
	findOptions := options.Find()
	var results []*Rubmats
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		println("Find user error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem Rubmats
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}
