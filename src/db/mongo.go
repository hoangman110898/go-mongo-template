package db

import (
	"go-mongo-template/src/config"
	"log"

	"gopkg.in/mgo.v2"
)

var instance *mgo.Session

var err error

func GetInstance(c *config.Configuration) *mgo.Session {
	if instance == nil {
		instance, err = mgo.Dial(c.MongoUrl)
		if err != nil {
			log.Printf("Mongo connect failed")
			panic(err)
		}
		log.Println("😺 mongo connect success 😋!")
	}
	return instance.Copy()
}
