package storage

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type MongoDBDriver struct {
	Session *mgo.Session
}

func NewMongoDriver(url string) (*MongoDBDriver, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	md := &MongoDBDriver{
		Session: s,
	}

	return md, nil
}

func (d *MongoDBDriver) Healthy() error {
	return d.Session.Ping()
}

func millisecondNow() int64 {
	s := time.Now()
	return s.Unix()*1000 + int64(s.Nanosecond()/1e6)
}
