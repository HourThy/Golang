package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	user "packages/User"
	login "packages/login"
	rubmats "packages/rubmats"
	"strconv"
	"time"
)

// User : User struct
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username, omitempty"`
	Password string             `bson:"password, omitempty"`
}

// Rubmats : Rubmate
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

// MongoLogin : To check user login isvalid
func MongoLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	result := login.CheckLogin(username, password)
	if result == "" {
		w.Write([]byte("Login successful!"))
	}
	w.Write([]byte("Wrong username or password!"))
}

// UnixLogin : To check user login isvalid
func UnixLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	result := login.FTPConnect(username, password)
	if result == "" {
		w.Write([]byte("Login successful!"))
	}
	w.Write([]byte("Wrong username or password!"))
}

// UnixLogin : To check user login isvalid
// func UnixLogin(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	result := unix.CheckUnixLogin(username, password)
// 	if result == "Success" {
// 		w.Write([]byte("Login successful!"))
// 	}
// 	w.Write([]byte("Login failed!"))
// }

// AddUser : To insert a new user
func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	result := user.InsertUser(username, password)
	w.Write([]byte(result))
}

// DeleteUser : Delete user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	user.DeleteUser(id)
}

// GetUser : Getting all Database to table in html
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:12345@198.1.1.95:27017/?authSource=dbLogin"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("dbLogin").Collection("User")
	findOptions := options.Find()
	var results []*User
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		println("Find user error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem User
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
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			panic("Wrong username or password!")
		}
		fmt.Printf("%s\n", b)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// UpdateUser : Update user information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	username := r.FormValue("username")
	password := r.FormValue("password")
	user.UpdateRecord(id, username, password)
}

// GetRubmats : GetRubmats
func GetRubmats(w http.ResponseWriter, r *http.Request) {
	// get value from input
	Machno := r.FormValue("findMachno")
	Matetype := r.FormValue("findMateType")
	Matecode := r.FormValue("findMateCode")
	Matename := r.FormValue("findMateName")

	//connect to mongo
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Database name and table name
	collection := client.Database("KV2MongoDB").Collection("rubmats")
	if err != nil {
		println("Connecing to MongoDB failed!")
	}
	fmt.Println("Connected to MongoDb!")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		println(err)
	}

	// Nếu machno là null thì tìm tất cả
	if Machno == "" && Matetype == "" && Matecode == "" && Matename == "" {
		findOptions := options.Find()
		var mulitipleResults []*Rubmats
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			println("Find machno error!")
		}
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				println(err)
			}
			mulitipleResults = append(mulitipleResults, &elem)
		}
		if err := cur.Err(); err != nil {
			println(err)
		}
		defer cur.Close(ctx)
		cur.Close(context.TODO())
		fmt.Printf("Found multiple documents (array of pointers): %+v\n", mulitipleResults)
		if r.Method == "GET" {
			b, err := json.MarshalIndent(mulitipleResults, "", "  ")
			if err != nil {
				panic("Return document to client error!")
			}
			fmt.Printf("%s\n", b)
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		}
		return
	}

	//Ngược lại nếu machno có liệu thì tìm machno trong bảng
	findOptions := options.Find()
	//findOptions.SetLimit(2)
	var result []*Rubmats
	// Finding multiple documents returns a cursor
	if Machno != "" {
		cur, err := collection.Find(context.TODO(), bson.D{{"machno", Machno}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, &elem)

		}

		if result == nil {
			w.Write([]byte("Error"))
			panic("No result")
		}
		// Close the cursor once finished
		cur.Close(context.TODO())
	} else if Matetype != "" {
		cur, err := collection.Find(context.TODO(), bson.D{{"matetype", Matetype}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, &elem)

		}

		if result == nil {
			w.Write([]byte("Error"))
			panic("No result")
		}
		// Close the cursor once finished
		cur.Close(context.TODO())
	} else if Matecode != "" {
		cur, err := collection.Find(context.TODO(), bson.D{{"matecode", Matecode}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, &elem)

		}

		if result == nil {
			w.Write([]byte("Error"))
			panic("No result")
		}
		// Close the cursor once finished
		cur.Close(context.TODO())
	} else if Matename != "" {
		cur, err := collection.Find(context.TODO(), bson.D{{"matename", Matename}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, &elem)

		}

		if result == nil {
			w.Write([]byte("Error"))
			panic("No result")
		}
		// Close the cursor once finished
		cur.Close(context.TODO())
	} else {
		panic("No result")
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	fmt.Println("Connectio to MongoDB close.")
	if r.Method == "GET" {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			panic("Return document to client error!")
		}
		fmt.Printf("%s\n", b)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

}

// GetAllData : GetAllData
// func GetAllData(w http.ResponseWriter, r *http.Request) {
// 	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	// Database name and table name
// 	collection := client.Database("KV2MongoDB").Collection("rubmats")
// 	if err != nil {
// 		println("Connecing to MongoDB failed!")
// 	}
// 	fmt.Println("Conected to MongoDb!")
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	findOptions := options.Find()
// 	var mulitipleResults []*Rubmats
// 	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
// 	if err != nil {
// 		println("Find machno error!")
// 	}
// 	for cur.Next(context.TODO()) {
// 		var elem Rubmats
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		mulitipleResults = append(mulitipleResults, &elem)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)
// 	fmt.Printf("Found multiple documents (array of pointers): %+v\n", mulitipleResults)
// 	if r.Method == "GET" {
// 		b, err := json.MarshalIndent(mulitipleResults, "", "  ")
// 		if err != nil {
// 			panic("Return document to client error!")
// 		}
// 		fmt.Printf("%s\n", b)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(b)
// 	}
// }

// InsertRubmats : InsertRubmats
func InsertRubmats(w http.ResponseWriter, r *http.Request) {
	machno := r.FormValue("machno")
	matetype := r.FormValue("matetype")
	matecode := r.FormValue("matecode")
	matename := r.FormValue("matename")
	intime := r.FormValue("intime")
	indat := r.FormValue("indat")
	usrno := r.FormValue("usrno")
	state := r.FormValue("state")
	if machno != "" && matetype != "" && matecode != "" && intime != "" && indat != "" && usrno != "" && state != "" {
		rubmats.InsertRubmats(machno, matetype, matecode, matename, intime, indat, usrno, state)
	}
	w.Write([]byte("Error to insert"))
}

// UpdateRubmats : UpdateRubmats
func UpdateRubmats(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("updateID")
	machno := r.FormValue("machno")
	matetype := r.FormValue("matetype")
	matecode := r.FormValue("matecode")
	matename := r.FormValue("matename")
	intime := r.FormValue("intime")
	indat := r.FormValue("indat")
	usrno := r.FormValue("usrno")
	state := r.FormValue("state")
	rubmats.UpdateRubmats(id, machno, matetype, matecode, matename, intime, indat, usrno, state)
}

// DeleteRubmats : DeleteRubmats
func DeleteRubmats(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	rubmats.DeleteRubmats(id)
}

// ServeHTTP : ServeHTTP
// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message": "hello world"}`))
// }

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")
	w.Write([]byte(fmt.Sprintf(`{"name": %d, "class": %d }`, userID, commentID, location)))
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	http.Handle("/register", http.StripPrefix("/register", http.FileServer(http.Dir("pages/register"))))
	http.Handle("/ulogin", http.StripPrefix("/ulogin", http.FileServer(http.Dir("pages/unix-login"))))
	http.Handle("/adduser", http.StripPrefix("/adduser", http.FileServer(http.Dir("pages/user_add"))))
	http.Handle("/getuser", http.StripPrefix("/getuser", http.FileServer(http.Dir("pages/user_get"))))
	http.Handle("/updateuser", http.StripPrefix("/updateuser", http.FileServer(http.Dir("pages/user_update"))))
	http.Handle("/home", http.StripPrefix("/home", http.FileServer(http.Dir("pages/home"))))
	http.Handle("/profile", http.StripPrefix("/profile", http.FileServer(http.Dir("pages/profile"))))
	http.Handle("/P9100", http.StripPrefix("/P9100", http.FileServer(http.Dir("pages/P9100"))))

	http.Handle("/example", http.StripPrefix("/example", http.FileServer(http.Dir("pages/example"))))
	http.HandleFunc("/mongoLogin", MongoLogin)
	http.HandleFunc("/unixLogin", UnixLogin)
	http.HandleFunc("/insertUser", AddUser)

	http.HandleFunc("/getuserlist", GetUser)
	http.HandleFunc("/deleteUser", DeleteUser)
	http.HandleFunc("/updateUser", UpdateUser)

	// Table rubmats
	http.Handle("/updaterubmats", http.StripPrefix("/updaterubmats", http.FileServer(http.Dir("pages/update_rubmats"))))
	http.HandleFunc("/getRubmats", GetRubmats)
	http.HandleFunc("/insertRubmats", InsertRubmats)
	http.HandleFunc("/updateRubmat", UpdateRubmats)
	http.HandleFunc("/deleteRubmats", DeleteRubmats)
	//http.HandleFunc("/getAllData", GetAllData)
	// Table rubmats

	// API
	//s := &server{}
	//http.Handle("/api", s)
	// r := mux.NewRouter()
	// api := r.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("/get", get).Methods(http.MethodGet)
	// api.HandleFunc("/post", post).Methods(http.MethodPost)
	// api.HandleFunc("/put", put).Methods(http.MethodPut)
	// api.HandleFunc("/delete", delete).Methods(http.MethodDelete)
	// api.HandleFunc("/notfound", notFound)
	// API
	log.Fatal(http.ListenAndServe(":5050", nil))
	//http.ListenAndServe(":5050", nil)
}
