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

	s.DB("").C(EntryCollection).DropCollection()

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
		desc        string
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
			err:  nil,
			desc: "Standart add entry",
			entryToAdd: &perfumepb.Entry{
				Id:     "2",
				Link:   "twitter.com",
				Title:  "Twitter second!",
				Labels: []string{"twitter", "golang"},
			},
		},
		{
			initialData: []*dbModel{},
			err:         nil,
			entryToAdd: &perfumepb.Entry{
				Link:   "facebook.com",
				Title:  "Facebook first!",
				Labels: []string{"facebook", "react"},
			},
			desc: "Add entry to empty db",
		},
		{
			initialData: []*dbModel{},
			err:         ErrorInvalidLink,
			entryToAdd: &perfumepb.Entry{
				Title:  "No link title!",
				Labels: []string{"testing", "purposes"},
			},
			desc: "Add entry without link given",
		},
		// TODO: Add other cases
	}

	for i, test := range tests {
		fmt.Printf("CASE %d: ", i)
		session := connect(test.initialData)

		err := session.AddEntry(test.entryToAdd)

		if err == test.err {
			fmt.Println("PASS")
		} else {
			fmt.Println(test.desc)
			t.Fatalf("Expected err: %v but got err: %v", test.err, err)
		}
	}
}

func TestGetEntry(t *testing.T) {
	tests := []struct {
		initialData []*dbModel
		desc        string
		err         error
		IDToGet     string
	}{
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc:    "Get existing entry with ID",
			err:     nil,
			IDToGet: "1",
		},
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc:    "Get non existing entry with ID",
			err:     ErrorNotFound,
			IDToGet: "3",
		},
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc:    "Give empty ID",
			err:     ErrorNotFound,
			IDToGet: "",
		},
	}

	for i, test := range tests {
		fmt.Printf("CASE %d:", i)
		session := connect(test.initialData)

		_, err := session.GetEntry(test.IDToGet)

		if err == test.err {
			fmt.Println("PASS")
		} else {
			fmt.Println(test.desc)
			t.Fatalf("Expected err: '%v' but got err: '%v'", err, test.err)
		}
	}
}

func TestUpdateEntry(t *testing.T) {
	tests := []struct {
		initialData   []*dbModel
		desc          string
		err           error
		EntryToUpdate *perfumepb.Entry
	}{
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc: "Update existing entry",
			err:  nil,
			EntryToUpdate: &perfumepb.Entry{
				Id:     "1",
				Title:  "Updated!",
				Link:   "http://twitter.com",
				Labels: []string{"twitter", "python", "data-science"},
			},
		},
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc: "Update non existing entry with ID",
			err:  ErrorNotFound,
			EntryToUpdate: &perfumepb.Entry{
				Id:     "3",
				Title:  "Updated!",
				Link:   "http://facebook.com",
				Labels: []string{"facebook", "react"},
			},
		},
	}

	for i, test := range tests {
		fmt.Printf("CASE %d:", i)
		session := connect(test.initialData)

		err := session.UpdateEntry(test.EntryToUpdate)

		if err == test.err {
			fmt.Println("PASS")
		} else {
			fmt.Println(test.desc)
			t.Fatalf("Expected err: '%v' but got err: '%v'", err, test.err)
		}
	}
}

func TestDeleteEntry(t *testing.T) {
	tests := []struct {
		initialData   []*dbModel
		desc          string
		err           error
		EntryToDelete *perfumepb.Entry
	}{
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc: "Delete non existing entry with ID",
			err:  ErrorNotFound,
			EntryToDelete: &perfumepb.Entry{
				Id: "3",
			},
		},
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc: "Delete existing entry",
			err:  nil,
			EntryToDelete: &perfumepb.Entry{
				Id: "1",
			},
		},
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc:          "Delete existing entry",
			err:           ErrorInvalidId,
			EntryToDelete: &perfumepb.Entry{},
		},
	}

	for i, test := range tests {
		fmt.Printf("CASE %d:", i)
		session := connect(test.initialData)

		err := session.DeleteEntry(test.EntryToDelete)

		if err == test.err {
			fmt.Println("PASS")
		} else {
			fmt.Println(test.desc)
			t.Fatalf("Expected err: '%v' but got err: '%v'", err, test.err)
		}
	}
}

func TestGetEntryList(t *testing.T) {
	tests := []struct {
		initialData []*dbModel
		desc        string
		err         error
	}{
		{
			initialData: []*dbModel{
				&dbModel{
					ID:     "1",
					Title:  "Twitter first!",
					Link:   "twitter.com",
					Labels: []string{"twitter", "python"},
				},
				&dbModel{
					ID:     "2",
					Title:  "Google second!",
					Link:   "google.com",
					Labels: []string{"google", "golang"},
				},
			},
			desc: "Get entry list",
			err:  nil,
		},
		{
			initialData: []*dbModel{},
			desc:        "Get empty entry list",
			err:         nil,
		},
	}

	for i, test := range tests {
		fmt.Printf("CASE %d:", i)
		session := connect(test.initialData)

		_, err := session.GetEntryList()

		if err == test.err {
			fmt.Println("PASS")
		} else {
			fmt.Println(test.desc)
			t.Fatalf("Expected err: '%v' but got err: '%v'", err, test.err)
		}
	}
}
