package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MDB *mongoDB = nil

type mongoDB struct {
	Mongo                 *mongo.Client
	machineConnCollection *mongo.Collection
	machineInfoCollection *mongo.Collection
	machineTMCollection   *mongo.Collection
}

func InitMongo(ctx context.Context, uri, db string, eas int64) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Log.Fatalf("Connect mongodb failed: %v", err)
		return err
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Log.Fatalf("Ping mongodb failed: %v", err)
		return err
	}
	MDB = &mongoDB{
		Mongo: client,
	}

	cl, err := client.Database(db).ListCollectionNames(ctx, bson.M{"name": "machine_tm"})
	if err != nil {
		log.Log.Fatalf("List mongodb collection names failed: %v", err)
		return err
	}
	if len(cl) == 0 {
		// Create collection with time series for machine info
		tsOpts := options.TimeSeries()
		tsOpts.SetTimeField("timestamp")
		tsOpts.SetMetaField("machine")
		tsOpts.SetGranularity("minutes")
		// tsOpts.SetBucketMaxSpan(30)
		// tsOpts.SetBucketRounding(5)
		ccOpts := options.CreateCollection()
		ccOpts.SetTimeSeriesOptions(tsOpts)
		ccOpts.SetExpireAfterSeconds(eas)
		if err := client.Database(db).CreateCollection(ctx, "machine_tm", ccOpts); err != nil {
			log.Log.Fatalf("Create time series collection failed: %v", err)
			return err
		}
		log.Log.Info("Create collection with time series success")
	}

	MDB.machineConnCollection = client.Database(db).Collection("machine_connection")
	MDB.machineInfoCollection = client.Database(db).Collection("machine_info")
	MDB.machineTMCollection = client.Database(db).Collection("machine_tm")
	return nil
}

func (db *mongoDB) Disconnect(ctx context.Context) error {
	return db.Mongo.Disconnect(ctx)
}

func (db *mongoDB) IsMachineConnected(ctx context.Context, machine types.MachineKey) bool {
	result := types.MDBMachineOnline{}
	if err := db.machineConnCollection.FindOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
	).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		return false
	}
	return true
}

func (db *mongoDB) MachineConnected(ctx context.Context, machine types.MachineKey) error {
	res, err := db.machineConnCollection.InsertOne(ctx, types.MDBMachineOnline{
		MachineKey: machine,
		AddTime:    time.Now(),
	})
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("insert online failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("inserted online id ", res.InsertedID)

	update, err := db.machineInfoCollection.UpdateOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
		bson.M{
			"$set": bson.M{
				"last_disconnect_time": time.Time{},
			},
		},
		// options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("reset last_disconnect_time failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("reset last_disconnect_time count ", update.ModifiedCount)
	return nil
}

func (db *mongoDB) MachineDisconnected(ctx context.Context, machine types.MachineKey) error {
	result, err := db.machineConnCollection.DeleteOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("delete online failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("delete online count ", result.DeletedCount)

	// update, err := db.machineInfoCollection.UpdateOne(
	// 	ctx,
	// 	bson.M{
	// 		"machine_id":   machine.MachineId,
	// 		"project":      machine.Project,
	// 		"container_id": machine.ContainerId,
	// 	},
	// 	bson.M{
	// 		"$set": bson.M{
	// 			"last_disconnect_time": time.Now(),
	// 		},
	// 	},
	// 	// options.Update().SetUpsert(true),
	// )
	// if err != nil {
	// 	log.Log.WithFields(logrus.Fields{"machine": machine}).Error("update last_disconnect_time failed: ", err)
	// 	return err
	// }
	// log.Log.WithFields(logrus.Fields{"machine": machine}).Info("update last_disconnect_time count ", update.ModifiedCount)
	return nil
}

func (db *mongoDB) GetMachineInfo(ctx context.Context, machine types.MachineKey) (*types.MDBMachineInfo, error) {
	result := &types.MDBMachineInfo{}
	if err := db.machineInfoCollection.FindOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
	).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *mongoDB) RegisterMachine(ctx context.Context, machine types.MachineKey, stakingType types.StakingType) error {
	res, err := db.machineInfoCollection.InsertOne(ctx, types.MDBMachineInfo{
		MachineKey:                      machine,
		StakingType:                     uint8(stakingType),
		RegisterTime:                    time.Now(),
		MDBDeepLinkMachineInfoST:        types.MDBDeepLinkMachineInfoST{},
		MDBDeepLinkMachineInfoBandwidth: types.MDBDeepLinkMachineInfoBandwidth{},
	})
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("insert machine info failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("inserted machine info id ", res.InsertedID)
	return nil
}

func (db *mongoDB) UnregisterMachine(ctx context.Context, machine types.MachineKey) error {
	result, err := db.machineInfoCollection.DeleteOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("delete machine info failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("delete machine info count ", result.DeletedCount)
	return nil
}

func (db *mongoDB) SetMachineInfo(
	ctx context.Context,
	machine types.MachineKey,
	deeplink_st *types.MDBDeepLinkMachineInfoST,
	deeplink_bw *types.MDBDeepLinkMachineInfoBandwidth,
) error {
	update, err := db.machineInfoCollection.UpdateOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
		bson.M{
			"$set": bson.M{
				"deeplink_st": deeplink_st,
				"deeplink_bw": deeplink_bw,
			},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("set machine info failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("set machine info count ", update.ModifiedCount)
	return nil
}

func (db *mongoDB) ReadDelayOffline(ctx context.Context, call func(types.MachineKey, time.Time, types.StakingType)) error {
	cur, err := db.machineInfoCollection.Find(ctx, bson.D{})
	if err != nil {
		return fmt.Errorf("get machine info iterater cursor failed: %v", err)
	}
	for cur.Next(ctx) {
		result := types.MDBMachineInfo{}
		if err := cur.Decode(&result); err != nil {
			return fmt.Errorf("decode machine info cursor into struct failed: %v", err)
		} else {
			// log.Log.Printf("cursor %v -> %v", result, cur.Current.Lookup("_id"))
			if !result.LastDisconnectTime.IsZero() {
				// call(result.MachineKey, time.Now(), types.StakingType(result.StakingType))
				call(result.MachineKey, time.Now(), types.StakingType(result.StakingType))
			}
		}
	}
	if err := cur.Err(); err != nil {
		return fmt.Errorf("machine info cursor error: %v", err)
	}
	cur.Close(ctx)
	return nil
}

func (db *mongoDB) WriteAllDelayOffline(ctx context.Context, tm time.Time) error {
	update, err := db.machineInfoCollection.UpdateMany(
		ctx,
		bson.M{},
		bson.M{
			"$set": bson.M{
				"last_disconnect_time": tm,
			},
		},
		// options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Log.Error("update all last_disconnect_time failed: ", err)
		return err
	}
	log.Log.Info("update all last_disconnect_time count ", update.ModifiedCount)
	return nil
}

func (db *mongoDB) OfflineMachine(ctx context.Context, machine types.MachineKey, tm time.Time) error {
	update, err := db.machineInfoCollection.UpdateOne(
		ctx,
		bson.M{
			"machine_id":   machine.MachineId,
			"project":      machine.Project,
			"container_id": machine.ContainerId,
		},
		bson.M{
			"$set": bson.M{
				"last_offline_time": tm,
			},
		},
		// options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": machine}).Error("update last_offline_time failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": machine}).Info("update last_offline_time count ", update.ModifiedCount)
	return nil
}

func (db *mongoDB) AddMachineTM(ctx context.Context, mtm types.MDBMachineTM) error {
	result, err := db.machineTMCollection.InsertOne(
		ctx,
		mtm,
	)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": mtm.Machine}).Error("insert machine tm info failed: ", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{"machine": mtm.Machine}).Info("inserted machine tm info id ", result.InsertedID)
	return nil
}

func (db *mongoDB) DeleteExpiredMachineTM(ctx context.Context, tm time.Time) error {
	result, err := db.machineTMCollection.DeleteMany(
		ctx,
		bson.M{
			"timestamp": bson.M{"$lt": tm},
		},
	)
	if err != nil {
		log.Log.Errorf("Delete expired documents before %v manully failed: %v", tm, err)
		return err
	}
	log.Log.Infof("Delete expired documents before %v manully DeletedCount %v", tm, result.DeletedCount)
	return nil
}

func (db *mongoDB) GetAllLatestMachineTM(ctx context.Context) []types.MDBMachineTM {
	di := make([]types.MDBMachineTM, 0)
	pipeline := mongo.Pipeline{
		// {{"$match", bson.D{{"timestamp", bson.D{{"$gt", specificTimestamp}}}}}},
		{{"$sort", bson.D{{"machine.machine_id", 1}, {"timestamp", -1}}}},
		{{"$group", bson.D{
			{"_id", "$machine.machine_id"},
			{"latestRecord", bson.D{{"$first", "$$ROOT"}}},
		}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$latestRecord"}}}},
	}
	cursor, err := db.machineTMCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Log.Errorf("Aggregate documents of all latest machine info failed: %v", err)
		return di
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := types.MDBMachineTM{}
		if err := cursor.Decode(&result); err != nil {
			log.Log.Errorf("Decode aggregate cursor into struct failed: %v", err)
		} else {
			di = append(di, result)
		}
	}
	if err := cursor.Err(); err != nil {
		log.Log.Errorf("Traversal aggregate cursor failed: %v", err)
	}
	return di
}
