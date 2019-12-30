package p1100

import (
	"config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	model "MES_Model"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/mgo.v2"
	bbson "gopkg.in/mgo.v2/bson"
)

// PostProCate : PostProCate
func PostProCate(w http.ResponseWriter, r *http.Request) {
	var count int
	if r.Method == "POST" {
		count++
		fmt.Println(count)
		valType := r.FormValue("subitem")
		arrVal := strings.Split(valType, "-")
		var code string
		var subitem string
		for i := 0; i < len(arrVal); i++ {
			code = arrVal[0]
			subitem = arrVal[1]
		}
		usr := r.FormValue("usr")
		indat := r.FormValue("indat")
		indat = strings.Replace(indat, "-", "", 64)
		intime := r.FormValue("intime")
		lot_id := code + subitem + indat + "000" + strconv.Itoa(count)
		fmt.Println(lot_id)
		fmt.Println(valType + "\t" + usr + "\t" + indat + "\t" + intime)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		collection := client.Database("KV2MongoDB").Collection("APRDCT")
		var results []*model.APRDCT

		projection := bson.D{
			{"PRODUCT_ID", 1},
			{"PRODUCT_DSC", 1},
			{"_id", 0},
		}

		cur, err := collection.Find(
			context.Background(),
			bson.D{
				{"PRODUCT_CATE", code},
			},
			options.Find().SetProjection(projection),
		)
		// cur, err := collection.Find(context.TODO(),
		// 	bson.D{{{PRODUCT_CATE: "Z"}, {PRODUCT_ID: 1, PRODUCT_DSC: 1}}}, findOptions)
		if err != nil {
			println("Find data error")
		}
		defer cur.Close(ctx)
		//fmt.Println(cur)
		for cur.Next(context.TODO()) {
			var elem model.APRDCT
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
			}
			results = append(results, &elem)
		}
		if err := cur.Err(); err != nil {
			panic(err)
		}
		//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// GetProCate : GetProCate
func GetProCate(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*model.ACODE
	cur, err := collection.Find(context.TODO(),
		// bson.D{{
		// 	"SUBITEM",
		// 	bson.D{{
		// 		"$in",
		// 		bson.A{"V", "D", "G", "RUBB", "Z"},
		// 	}},
		// }}, findOptions)
		bson.M{"CODE_CATE": "PCAT"}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem model.ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)

	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	//fmt.Println("method:", r.Metho)
	if r.Method == "GET" {
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

var prodID string
var prodCate string
var endDate string
var startDate string

// PostProID : PostProID
func PostProID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		prodID = r.FormValue("prodID")
		prodCate = r.FormValue("prodCate")
		endDate = r.FormValue("endDate")
		startDate = r.FormValue("startDate")
		db, err := config.GetMongoDB()
		if err != nil {
			panic(err)
		} else {
			var results []model.AEQPTRESV
			err2 := db.C("AEQPTRESV").Find(
				bson.M{
					"LOT_ID": bson.M{"$regex": bbson.RegEx{"^.*" + prodCate + ".*$", "i"}},
					"NX_OPE_ID": bson.M{"$regex": bbson.RegEx{"^.*" + prodID + ".*$", "i"}},
					"RESV_DATE": bson.M{
						"$gte": startDate,
						"$lte": endDate,
					}}).Select(
				bson.M{
					"LOT_ID":          1,
					"NX_OPE_ID":       1,
					"RESV_EQPT_ID":    1,
					"LOT_STAT":        1,
					"RESV_DATE":       1,
					"RESV_SHIFT_SEQ":  1,
					"RESV_COMMENT":    1,
					"CLAIM_DATE":      1,
					"CLAIM_TIME":      1,
					"CLAIM_USER":      1,
					"PLAN_OPT_WEIGHT": 1,
					"SHT_CNT":         1,
				}).All(&results)
			if err2 != nil {
				panic(err)
			} else {
				b, err := json.MarshalIndent(results, "", "\t")
				if err != nil {
					panic("Err!")
				}
				//fmt.Printf("%s\n", )
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
			}
		}
	}
}

// GetLineID : GetLineID
func GetLineID(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*model.ACODE
	cur, err := collection.Find(context.TODO(),
		// bson.D{{
		// 	"SUBITEM",
		// 	bson.D{{
		// 		"$in",
		// 		bson.A{"V", "D", "G", "RUBB", "Z"},
		// 	}},
		// }}, findOptions)
		bson.M{"CODE_CATE": "LNID"}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem model.ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		//strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)

	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	//fmt.Println("method:", r.Metho)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// PostLineID : PostLineID
// func PostLineID(w http.ResponseWriter, r *http.Request) {
// 	var count int
// 	if r.Method == "POST" {
// 		count++
// 		valType := r.FormValue("subitem")
// 		arrVal := strings.Split(valType, "-")
// 		var code string
// 		var subitem string
// 		for i := 0; i < len(arrVal); i++ {
// 			code = arrVal[0]
// 			subitem = arrVal[1]
// 		}
// 		usr := r.FormValue("usr")
// 		indat := r.FormValue("indat")
// 		indat = strings.Replace(indat, "-", "", 64)
// 		intime := r.FormValue("intime")
// 		lot_id := code + subitem + indat + "000" + strconv.Itoa(count)
// 		fmt.Println(lot_id)
// 		fmt.Println(valType + "\t" + usr + "\t" + indat + "\t" + intime)
// 		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = client.Ping(context.TODO(), nil)
// 		if err != nil {
// 			panic(err)
// 		}
// 		collection := client.Database("KV2MongoDB").Collection("APRDCT")
// 		//findOptions := options.Find()
// 		var results []*model.APRDCT

// 		projection := bson.D{
// 			{"PRODUCT_ID", 1},
// 			{"PRODUCT_DSC", 1},
// 			{"_id", 0},
// 		}

// 		cur, err := collection.Find(
// 			context.Background(),
// 			bson.D{
// 				{"PRODUCT_CATE", code},
// 			},
// 			options.Find().SetProjection(projection),
// 		)
// 		// cur, err := collection.Find(context.TODO(),
// 		// 	bson.D{{{PRODUCT_CATE: "Z"}, {PRODUCT_ID: 1, PRODUCT_DSC: 1}}}, findOptions)
// 		if err != nil {
// 			println("Find data error")
// 		}
// 		defer cur.Close(ctx)
// 		//fmt.Println(cur)
// 		for cur.Next(context.TODO()) {
// 			var elem model.APRDCT
// 			err := cur.Decode(&elem)
// 			if err != nil {
// 				panic(err)
// 			}
// 			results = append(results, &elem)
// 		}
// 		if err := cur.Err(); err != nil {
// 			panic(err)
// 		}
// 		//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
// 		b, err := json.MarshalIndent(results, "", "\t")
// 		if err != nil {
// 			panic("Err!")
// 		}
// 		//fmt.Printf("%s\n", b)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(b)
// 	}
// }

// GetOwnerID : GetOwnerID
func GetOwnerID(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*model.ACODE
	cur, err := collection.Find(context.TODO(),
		bson.M{"CODE_CATE": "OWNR"}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem model.ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		//strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", b)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// GetMaxLOTID : GetMaxLOTID
func GetMaxLOTID(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("KV2MongoDB").C("AEQPTRESV")
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"NX_OPE_ID": prodID, "LOT_ID": bbson.RegEx{Pattern: "^" + prodCate, Options: "i"}, "RESV_DATE": bson.M{"$gte": startDate, "$lte": endDate}}},
		bson.M{"$group": bson.M{"_id": "", "MAX(LOT_ID)": bson.M{"$max": "$LOT_ID"}}},
		bson.M{"$project": bson.M{"MAX(LOT_ID)": "$MAX(LOT_ID)", "_id": 0}},
	}
	pipe := collection.Pipe(pipeline)
	resp := []bson.M{}
	err = pipe.All(&resp)
	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}
	b, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// InsertAEQPTRESV : InsertAEQPTRESV
func InsertAEQPTRESV(w http.ResponseWriter, r *http.Request) {
	LOT_ID := r.FormValue("LOT_ID")
	//check_LOT_ID := strings.substring
	NX_OPE_NO := "";
	NX_OPE_ID := r.FormValue("NX_OPE_ID")
	NX_OPE_VER := ""
	SPLT_ID := ""
	RESV_EQPT_ID := ""
	LOT_STAT := ""
	RUN_FLAG := ""
	RESV_DATE := r.FormValue("RESV_DATE")
	RESV_SHIFT_SEQ := ""
	RESV_COMMENT := ""
	CLAIM_DATE := r.FormValue("CLAIM_DATE")
	CLAIM_TIME := r.FormValue("CLAIM_TIME")
	CLAIM_USER := r.FormValue("CLAIM_USER")
	MOVE_IN_WEIGHT := ""
	MOVE_OUT_WEIGHT := ""
	PLAN_OPT_WEIGHT := r.FormValue("PLAN_OPT_WEIGHT")
	SHT_CNT := ""
	CR_SHT_CNT := ""
	FIT_EQPTS := ""
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		println(err)
	}

	collection := client.Database("KV2MongoDB").Collection("AEQPTRESV")
	filter := bson.D{
		{"LOT_ID", LOT_ID},
		{"NX_OPE_NO", NX_OPE_NO},
		{"NX_OPE_ID", NX_OPE_ID},
		{"NX_OPE_VER", NX_OPE_VER},
		{"SPLT_ID", SPLT_ID},
		{"RESV_EQPT_ID", RESV_EQPT_ID},
		{"LOT_STAT", LOT_STAT},
		{"RUN_FLAG", RUN_FLAG},
		{"RESV_DATE", RESV_DATE},
		{"RESV_SHIFT_SEQ", RESV_SHIFT_SEQ},
		{"RESV_COMMENT", RESV_COMMENT},
		{"CLAIM_DATE", CLAIM_DATE},
		{"CLAIM_TIME", CLAIM_TIME},
		{"CLAIM_USER", CLAIM_USER},
		{"MOVE_IN_WEIGHT", MOVE_IN_WEIGHT},
		{"MOVE_IN_WEIGHT", MOVE_OUT_WEIGHT},
		{"PLAN_OPT_WEIGHT", PLAN_OPT_WEIGHT},
		{"SHT_CNT", SHT_CNT},
		{"CR_SHT_CNT", CR_SHT_CNT},
		{"FIT_EQPTS", FIT_EQPTS},
	}
	insertResult, err := collection.InsertOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted many document: ", insertResult.InsertedID)
}
