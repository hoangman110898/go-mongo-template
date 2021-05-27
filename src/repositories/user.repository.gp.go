package repositories

import (
	"go-mongo-template/src/config"
	"go-mongo-template/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository interface {
	Create(*models.User) error
	Update(interface{}, interface{}) error
	Detail(string) error
}

type UserRepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func (repository *UserRepositoryImp) collection() *mgo.Collection {
	return repository.db.DB(repository.config.DatabaseName).C("users")
}

func (repository *UserRepositoryImp) Create(user *models.User) error {
	return repository.collection().Insert(user)
}

func (repository *UserRepositoryImp) Update(query, change interface{}) error {
	return repository.collection().Update(query, change)
}

func (repository *UserRepositoryImp) Detail(id string) error {
	var user models.User
	return repository.collection().Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&user)
}

func NewUser(db *mgo.Session, c *config.Configuration) UserRepository {
	return &UserRepositoryImp{db: db, config: c}
}
