package login

import (
	"context"
	"fmt"

	"github.com/jlaffaye/ftp"
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

// CheckLogin : CheckLogin function
func CheckLogin(username, password string) string {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoD")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("dbLogin").Collection("User")

	filter := bson.D{{"username", username}, {"password", password}}
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	if result.Username != username && result.Password != password {
		panic(err)
	}
	return username
}

// FTPConnect : Connect to FTP Server
func FTPConnect(username, password string) string {
	client, err := ftp.Dial("192.1.1.1:21")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer client.Quit()

	if err := client.Login(username, password); err != nil {
		panic(err)
	}
	return username
}
