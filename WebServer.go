package main

import (
	model "MES_Model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	login "packages/login"
	p1100 "packages/p1100"
	p1200 "packages/p1200"
	rubmats "packages/rubmats"
	"strings"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/goracle.v2"
)

// Users : Users struct
type Users struct {
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
// func AddUser(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	result := users.InsertUser(username, password)
// 	w.Write([]byte(result))
// }

// DeleteUser : Delete user from database
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	id := r.FormValue("id")
// 	users.DeleteUser(id)
// }

// GetUser : Getting all Database to table in html
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:12345@198.1.1.95:27017/?authSource=dbLogin"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("dbLogin").Collection("User")
	findOptions := options.Find()
	var results []*Users
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		println("Find user error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem Users
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		panic(err)
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
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	id := r.FormValue("id")
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	users.UpdateRecord(id, username, password)
// }

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
			panic(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
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
			panic(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
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
			panic(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
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
			panic(err)
		}
		// Iterate through the cursor
		for cur.Next(context.TODO()) {
			var elem Rubmats
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
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
// 		panic(err)
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
// 			panic(err)
// 		}
// 		mulitipleResults = append(mulitipleResults, &elem)
// 	}
// 	if err := cur.Err(); err != nil {
// 		panic(err)
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

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "123456"
	dbname   = "dbuser"
)

//Account : Account
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type allUsers []Account

// ConnectPostgreSQL : ConnectPostgreSQL
func ConnectPostgreSQL(w http.ResponseWriter, r *http.Request) {
	//time := time.NewTimer(10 * time.Second)
	//reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	// }

	psqlInfo := fmt.Sprintf("host=localhost port=8080 user=postgres " +
		"password=123456 dbname=dbuser sslmode=disable")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT id, username, password FROM tbusers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := []Account{}
	for rows.Next() {
		elem := Account{}
		err = rows.Scan(&elem.ID, &elem.Username, &elem.Password)
		if err != nil {
			panic(err)
		}
		result = append(result, elem)
	}
	e, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(e)
	//w.WriteHeader(http.StatusOK)
}

// ConnectOracle : ConnectOracle
func ConnectOracle(w http.ResponseWriter, r *http.Request) {
	// Oracle
	db, err := sql.Open("goracle", "system/password@198.1.9.189:1521/KVDBP01")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to Oracle!")
	defer db.Close()
	// Oracle

	// MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		println(err)
	}
	// MongoDB

	rows, err := db.Query("SELECT* FROM PPT1D.AEQPTRESV")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	collection := client.Database("KV2MongoDB").Collection("AEQPTRESV")
	var data []interface{}
	for rows.Next() {
		var elem model.AEQPTRESV
		if err = rows.Scan(
			&elem.LOT_ID,
			&elem.NX_OPE_NO,
			&elem.NX_OPE_ID,
			&elem.NX_OPE_VER,
			&elem.SPLT_ID,
			&elem.RESV_EQPT_ID,
			&elem.LOT_STAT,
			&elem.RUN_FLAG,
			&elem.RESV_DATE,
			&elem.RESV_SHIFT_SEQ,
			&elem.RESV_COMMENT,
			&elem.CLAIM_DATE,
			&elem.CLAIM_TIME,
			&elem.CLAIM_USER,
			&elem.MOVE_IN_WEIGHT,
			&elem.MOVE_OUT_WEIGHT,
			&elem.PLAN_OPT_WEIGHT,
			&elem.SHT_CNT,
			&elem.CR_SHT_CNT,
			&elem.FIT_EQPTS,
		); err != nil {
			panic(err)
		}
		//results = append(results, &elem)
		filter := bson.M{
			"LOT_ID":          strings.TrimSpace(elem.LOT_ID),
			"NX_OPE_NO":       strings.TrimSpace(elem.NX_OPE_NO),
			"NX_OPE_ID":       strings.TrimSpace(elem.NX_OPE_ID),
			"NX_OPE_VER":      strings.TrimSpace(elem.NX_OPE_VER),
			"SPLT_ID":         strings.TrimSpace(elem.SPLT_ID),
			"RESV_EQPT_ID":    strings.TrimSpace(elem.RESV_EQPT_ID),
			"LOT_STAT":        strings.TrimSpace(elem.LOT_STAT),
			"RUN_FLAG":        strings.TrimSpace(elem.RUN_FLAG),
			"RESV_DATE":       strings.TrimSpace(elem.RESV_DATE),
			"RESV_SHIFT_SEQ":  strings.TrimSpace(elem.RESV_SHIFT_SEQ),
			"RESV_COMMENT":    strings.TrimSpace(elem.RESV_COMMENT),
			"CLAIM_DATE":      strings.TrimSpace(elem.CLAIM_DATE),
			"CLAIM_TIME":      strings.TrimSpace(elem.CLAIM_TIME),
			"CLAIM_USER":      strings.TrimSpace(elem.CLAIM_USER),
			"MOVE_IN_WEIGHT":  strings.TrimSpace(elem.FIT_EQPTS),
			"MOVE_OUT_WEIGHT": elem.MOVE_OUT_WEIGHT,
			"PLAN_OPT_WEIGHT": elem.PLAN_OPT_WEIGHT,
			"SHT_CNT":         elem.SHT_CNT,
			"CR_SHT_CNT":      elem.CR_SHT_CNT,
			"FIT_EQPTS":       elem.FIT_EQPTS,
		}
		data = []interface{}{filter}
		insertManyResult, err := collection.InsertMany(context.TODO(), data)
		if err != nil {
			panic(err)
		}
		fmt.Println("Document inserted: ", insertManyResult.InsertedIDs)
	}
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
	http.HandleFunc("/getuserlist", GetUser)

	// Table rubmats
	http.Handle("/updaterubmats", http.StripPrefix("/updaterubmats", http.FileServer(http.Dir("pages/update_rubmats"))))
	http.HandleFunc("/getRubmats", GetRubmats)
	http.HandleFunc("/insertRubmats", InsertRubmats)
	http.HandleFunc("/updateRubmat", UpdateRubmats)
	http.HandleFunc("/deleteRubmats", DeleteRubmats)

	http.Handle("/postgre", http.StripPrefix("/postgre", http.FileServer(http.Dir("pages/postgre"))))
	http.Handle("/1200", http.StripPrefix("/1200", http.FileServer(http.Dir("pages/1200"))))
	http.HandleFunc("/connectToPostgreSQL", ConnectPostgreSQL)

	http.HandleFunc("/connectOracle", ConnectOracle)

	//Restful API
	// r := mux.NewRouter()
	// r.HandleFunc("/connectToPostgreSQL", ConnectPostgreSQL).Methods(http.MethodGet)
	// r.HandleFunc("/", post).Methods(http.MethodPost)
	// r.HandleFunc("/", put).Methods(http.MethodPut)
	// r.HandleFunc("/", delete).Methods(http.MethodDelete)
	// r.HandleFunc("/", notFound)

	// ********* Call Oracle ********
	http.Handle("/1100", http.StripPrefix("/1100", http.FileServer(http.Dir("pages/1100"))))
	http.HandleFunc("/getProCate", p1100.GetProCate)
	// http.HandleFunc("/getProID", GetProID)
	http.HandleFunc("/getLineID", p1100.GetLineID)
	http.HandleFunc("/getOwnerID", p1100.GetOwnerID)
	http.HandleFunc("/postProCate", p1100.PostProCate)
	http.HandleFunc("/getMaxLOTID", p1100.GetMaxLOTID)
	http.HandleFunc("/postProID", p1100.PostProID)
	http.HandleFunc("/insertAEQPTRESV", p1100.InsertAEQPTRESV)
	http.HandleFunc("/getBayID", p1200.GetBayID)
	http.HandleFunc("/postBayID", p1200.PostBayID)
	http.HandleFunc("/getMachineID", p1200.GetMachineID)
	// ********* Call Oracle ********

	http.ListenAndServe(":6060", nil)
}
