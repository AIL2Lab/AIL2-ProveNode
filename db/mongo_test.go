package db

import (
	"context"
	"errors"
	"log"
	"strconv"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// go test -v -timeout 30s -count=1 -run TestMongoDBBasic AIL2-ProveNode/db
func TestMongoDBBasic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var uri = "mongodb://localhost:27017"
	// var uri = "mongodb://localhost:37017"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		t.Fatalf("Connect mongodb failed: %v", err)
	}
	t.Log("Connect mongodb success")
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	searchCollection := func(ctx context.Context, client *mongo.Client) {
		cl, err := client.Database("test").ListCollectionNames(ctx, bson.M{"name": "testinfo", "type": "collection"})
		t.Logf("ListCollectionNames testinfo return %v %v", cl, err)
	}
	searchCollection(ctx, client)

	collection := client.Database("test").Collection("testinfo")
	for i := 0; i < 10; i++ {
		res, err := collection.InsertOne(ctx, bson.M{
			"name":        strconv.Itoa(i),
			"score":       i * 10,
			"add_time":    time.Now(), // time.Now().UnixMilli(),
			"update_time": time.Now(), // time.Now().UnixMilli(),
		})
		if err != nil {
			t.Fatalf("Insert mongodb item failed: %v", err)
		}
		t.Logf("Inserted id %v", res.InsertedID)
	}

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Name  string
			Score int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.Lookup("_id"))
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	result := struct {
		Name       string
		Score      int
		AddTime    time.Time `json:"add_time" bson:"add_time"`
		UpdateTime time.Time `json:"update_time" bson:"update_time"`
	}{}
	if err := collection.FindOne(ctx, bson.M{"name": "6"}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			t.Log("Find nothing in collection")
		} else {
			t.Fatalf("Find collection item failed: %v", err)
		}
	}
	t.Logf("Find collection item: %v", result)

	pipeline := mongo.Pipeline{
		{{"$sort", bson.D{{"score", -1}}}},
		{{"$limit", 1}},
	}
	cur, err = collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalf("Aggregate failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Name       string
			Score      int
			AddTime    time.Time `json:"add_time" bson:"add_time"`
			UpdateTime time.Time `json:"update_time" bson:"update_time"`
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)
	t.Log("Aggregate end")

	deleteRes, err := collection.DeleteOne(ctx, bson.M{"name": "6"})
	if err != nil {
		t.Fatalf("Delete collection item failed: %v", err)
	}
	t.Logf("Delete DeletedCount %v", deleteRes.DeletedCount)

	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Name       string
			Score      int
			Grade      int
			AddTime    time.Time `json:"add_time" bson:"add_time"`
			UpdateTime time.Time `json:"update_time" bson:"update_time"`
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.Lookup("_id"))
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	updateRes, err := collection.UpdateOne(ctx, bson.M{"name": "6"}, bson.M{
		"$set": bson.M{
			"score":       6,
			"grade":       12,
			"update_time": time.Now(),
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		t.Fatalf("Update collection item failed: %v", err)
	}
	t.Logf("Update MatchedCount %v ModifiedCount %v, UpsertedCount %v",
		updateRes.MatchedCount, updateRes.ModifiedCount, updateRes.UpsertedCount)

	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Name  string
			Score int
			Grade int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	searchCollection(ctx, client)

	time.Sleep(2 * time.Second)

	if err := collection.Drop(ctx); err != nil {
		t.Fatalf("Drop test database failed: %v", err)
	}
	t.Log("Drop test database success")

	searchCollection(ctx, client)
}

// go test -v -timeout 300s -count=1 -run TestMongoDBTimeSeries AIL2-ProveNode/db
func TestMongoDBTimeSeries(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var uri = "mongodb://localhost:27017"
	// var uri = "mongodb://localhost:37017"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		t.Fatalf("Connect mongodb failed: %v", err)
	}
	t.Log("Connect mongodb success")
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	searchCollection := func(ctx context.Context, client *mongo.Client) {
		cl, err := client.Database("test").ListCollectionNames(ctx, bson.M{"name": "testinfo"})
		t.Logf("ListCollectionNames testinfo return %v %v", cl, err)
	}
	searchCollection(ctx, client)

	tsOpts := options.TimeSeries()
	tsOpts.SetTimeField("timestamp")
	tsOpts.SetMetaField("device")
	tsOpts.SetGranularity("seconds")
	// tsOpts.SetBucketMaxSpan(30)
	// tsOpts.SetBucketRounding(5)
	ccOpts := options.CreateCollection()
	ccOpts.SetTimeSeriesOptions(tsOpts)
	ccOpts.SetExpireAfterSeconds(5)
	if err := client.Database("test").CreateCollection(ctx, "testinfo", ccOpts); err != nil {
		t.Fatalf("Create time series collection failed: %v", err)
	}

	searchCollection(ctx, client)

	collection := client.Database("test").Collection("testinfo")
	tm := time.Now()

	// bson.TypeDateTime
	docs := []interface{}{
		bson.M{
			"timestamp":   tm.Add(-90 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 15,
			"utilization": 30,
		},
		bson.M{
			"timestamp":   tm.Add(-60 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 16,
			"utilization": 50,
		},
		bson.M{
			"timestamp":   tm.Add(-30 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 12,
			"utilization": 25,
		},
		bson.M{
			"timestamp":   tm.Add(30 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 11,
			"utilization": 24,
		},
		bson.M{
			"timestamp":   tm.Add(60 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 13,
			"utilization": 28,
		},
		bson.M{
			"timestamp":   tm.Add(90 * time.Second),
			"device":      bson.M{"id": "5578", "project": "degpt"},
			"temperature": 12,
			"utilization": 25,
		},
		bson.M{
			"timestamp":   tm.Add(-90 * time.Second),
			"device":      bson.M{"id": "9527", "project": "degpt"},
			"temperature": 16,
			"utilization": 60,
		},
	}
	res, err := collection.InsertMany(ctx, docs, options.InsertMany().SetOrdered(false))
	if err != nil {
		t.Fatalf("Insert mongodb many failed: %v", err)
	}
	t.Logf("Inserted %v", res.InsertedIDs...)

	result := struct {
		Temperature int
		Utilization int
	}{}
	if err := collection.FindOne(ctx, bson.M{"timestamp": tm.Add(-60 * time.Second), "device.id": "5578"}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			t.Log("Find nothing in collection")
		} else {
			t.Fatalf("Find collection item failed: %v", err)
		}
	}
	t.Logf("Find collection item: %v", result)

	pipeline := mongo.Pipeline{
		{{"$sort", bson.D{{"device.id", 1}, {"timestamp", -1}}}},
		{{"$group", bson.D{
			{"_id", "$device.id"},
			{"latestRecord", bson.D{{"$first", "$$ROOT"}}},
		}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$latestRecord"}}}},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalf("Aggregate failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Timestamp time.Time
			Device    struct {
				Id      string
				Project string
			}
			Temperature int
			Utilization int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)
	t.Log("Aggregate end")

	time.Sleep(2 * time.Second)

	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Timestamp time.Time
			Device    struct {
				Id      string
				Project string
			}
			Temperature int
			Utilization int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	t.Log("Wait for deletion of expired document")
	time.Sleep(10 * time.Second)

	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Timestamp time.Time
			Device    struct {
				Id      string
				Project string
			}
			Temperature int
			Utilization int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	// delOpts := options.Delete().SetHint(bson.M{"device.id": "5578", "device.project": "degpt"})
	// delOpts := options.Delete().SetHint(bson.M{"device": bson.M{"id": "5578", "project": "degpt"}})
	deleteRes, err := collection.DeleteMany(
		ctx,
		bson.M{
			// "device.id":      "5578",
			// "device.project": "degpt",
			// "device":    bson.M{"id": "5578", "project": "degpt"},
			"timestamp": bson.M{"$lt": tm},
		},
		// delOpts,
	)
	if err != nil {
		t.Logf("Delete expired documents manully failed: %v", err)
	}
	t.Log("Delete expired documents manully with DeletedCount: ", deleteRes.DeletedCount)

	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Get iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := struct {
			Timestamp time.Time
			Device    struct {
				Id      string
				Project string
			}
			Temperature int
			Utilization int
		}{}
		if err := cur.Decode(&result); err != nil {
			t.Fatalf("decode cursor into struct failed: %v", err)
		}
		t.Logf("cursor %v -> %v", result, cur.Current.String())
	}
	if err := cur.Err(); err != nil {
		t.Fatalf("cursor error: %v", err)
	}
	cur.Close(ctx)

	if err := collection.Drop(ctx); err != nil {
		t.Fatalf("Drop test database failed: %v", err)
	}
	t.Log("Drop test database success")
}
