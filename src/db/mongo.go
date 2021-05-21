package db

import (
	"go-mongo-template/src/config"
	"gopkg.in/mgo.v2"
	"log"
)

var instance *mgo.Session

var err error

func GetInstance(c *config.Configuration) *mgo.Session  {
	if instance == nil {
		instance, err = mgo.Dial(c.MongoUrl)

		if err != nil {
			panic(err)
		}

		log.Fatalln("😺 mongo connect success 😋!")
	}

	return instance.Copy()
}

