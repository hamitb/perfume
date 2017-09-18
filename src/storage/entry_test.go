package storage

import (
	"fmt"
	"os"
	"perfumepb"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

type dbModel struct {
	ID        string   `bson:"_id"`
	CreatedAt int64    `bson:"created_at"`
	Link      string   `bson:"link"`
	Title     string   `bson:"title"`
	Labels    []string `bson:"labels"`
}

func conv(data []*dbModel) []interface{} {
	d := make([]interface{}, len(data))
	for i, v := range data {
		d[i] = v
	}
	return d
}

func connect(data []*dbModel) *MongoDBDriver {
	url := os.Getenv("TEST_DB_URL")
	if url == "" {
		url = "mongodb://localhost:27017"
	}
	s, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		if err := s.DB("").C(EntryCollection).Insert(conv(data)...); err != nil {
			panic(err)
		}
	}

	return &MongoDBDriver{
		Session: s,
	}
}

func TestAddEntry(t *testing.T) {
	tests := []struct {
		initialData []*dbModel
		entryToAdd  *perfumepb.Entry
		err         error
	}{
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Link:   "google.com",
					Title:  "Google first!",
					Labels: []string{"test", "aLabel"},
				},
			},
			err: nil,
			entryToAdd: &perfumepb.Entry{
				Id:     "2",
				Link:   "twitter.com",
				Title:  "Twitter second!",
				Labels: []string{"twitter", "golang"},
			},
		},
		// TODO: Add other cases
	}

	for i, test := range tests {
		fmt.Printf("Case %d: ", i)
		session := connect(test.initialData)

		err := session.AddEntry(test.entryToAdd)

		if err == nil && test.err == nil {
			fmt.Println("No error")
		}
	}
}
