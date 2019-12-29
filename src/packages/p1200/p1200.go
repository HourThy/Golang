package p1200

import (
	"config"
	"context"
	"encoding/json"
	"net/http"
	"time"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/goracle.v2"
	//"gopkg.in/mgo.v2"
	bbson "gopkg.in/mgo.v2/bson"
	model "MES_Model"
	"strings"
)

// // ABAYAREA : ABAYAREA struct
// type ABAYAREA struct {
// 	BAY_ID      string
// 	OPI_FLG     string
// 	BAY_DSC     string
// 	ADDT_INFO_1 string
// 	ADDT_INFO_2 string
// 	ADDT_INFO_3 string
// }

// GetBayID : GetBayID
func GetBayID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		collection := client.Database("KV2MongoDB").Collection("ABAYAREA")
		findOptions := options.Find()
		var results []*model.ABAYAREA
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			println("Find data error")
		}
		defer cur.Close(ctx)
		for cur.Next(context.TODO()) {
			var elem model.ABAYAREA
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
		//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
		//fmt.Println("method:", r.Metho)
		
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(b)
		
	}
}

// PostBayID : PostBayID
func PostBayID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		BAY_ID := r.FormValue("BAY_ID")
		BAY_ID = strings.TrimSpace(BAY_ID)
		if BAY_ID != "" {
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
			if err != nil {
				panic(err)
			}
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				panic(err)
			}
			collection := client.Database("KV2MongoDB").Collection("AEQPT")
			findOptions := options.Find()
			var results []*model.AEQPT
			cur, err := collection.Find(context.TODO(), bson.M{"BAY_ID": BAY_ID}, findOptions)
			if err != nil {
				println("Find data error")
			}
			defer cur.Close(ctx)
			for cur.Next(context.TODO()) {
				var elem model.AEQPT
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
			b, err := json.MarshalIndent(results, "", "\t")
			if err != nil {
				panic("Err!")
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Write(b)
		} else {
			w.Write([]byte(" Data is empty"))
		}
	}
}

// GetMachineID : GetMachineID
func GetMachineID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bayID := r.FormValue("BAY_ID")
		db, err := config.GetMongoDB()
		if err != nil {
			panic(err)
		} else {
			var results []model.AEQPTRESV
			err2 := db.C("AEQPTRESV").Find(bson.M{"FIT_EQPTS": bson.M{"$regex": bbson.RegEx{"^.*" + bayID + ".*$", "i"}}}).Select(
				bson.M{
					"LOT_ID": 1, 
					"NX_OPE_ID": 1,
					"NX_OPE_VER": 1,
					"PLAN_OPT_WEIGHT": 1,
					"NX_OPE_NO": 1,
					"RESV_DATE": 1,
					"RESV_SHIFT_SEQ": 1,
				},
				).All(&results)
			if err2 != nil {
				panic(err)
			} else {
				b, err := json.MarshalIndent(results, "", "\t")
				if err != nil {
					panic("Err!")
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
			}
		}
	}
}
