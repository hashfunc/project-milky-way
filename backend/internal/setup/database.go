package setup

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hashfunc/project-milky-way/internal"
	"github.com/hashfunc/project-milky-way/internal/config"
	"github.com/hashfunc/project-milky-way/internal/database"
)

func Database(stars []*database.Star) error {
	conf, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal(conf)
	}

	connection, err := database.New(&conf.Database)

	if err != nil {
		log.Fatal(err)
	}

	collection := connection.DB().Collection("stars")

	if collection == nil {
		log.Fatal("Collection 조회 실패")
	}

	view := collection.Indexes()

	exist, err := isIndexExist(&view)

	if err != nil {
		log.Fatal(err)
	}

	if !exist {
		if err := createIndex(collection); err != nil {
			log.Fatal(err)
		}
	}

	documents := make([]interface{}, len(stars))
	for index, star := range stars {
		documents[index] = star
	}

	_, err = collection.InsertMany(context.TODO(), documents)

	return err
}

func isIndexExist(view *mongo.IndexView) (bool, error) {
	cur, err := view.List(nil)

	if err != nil {
		return false, err
	}

	defer func() {
		if err := cur.Close(nil); err != nil {
			log.Panic(
				internal.NewError("인덱스 커서닫기 오류", err),
			)
		}
	}()

	for cur.Next(nil) {
		raw, err := cur.Current.LookupErr("key")

		if err != nil {
			log.Print(err)
			continue
		}

		key := map[string]interface{}{}
		if err := raw.Unmarshal(&key); err != nil {
			log.Print(err)
			continue
		}

		if key["point"] == "2dsphere" {
			return true, nil
		}
	}

	return false, nil
}

func createIndex(collection *mongo.Collection) error {
	unique := true

	index := []mongo.IndexModel{
		{
			Keys: bson.D{{"point", "2dsphere"}},
		},
		{
			Keys:    bson.D{{"code", 1}},
			Options: &options.IndexOptions{Unique: &unique},
		},
	}

	_, err := collection.Indexes().CreateMany(context.TODO(), index)

	return err
}
