package repositories

import (
	"go-mongo-template/src/config"
	"go-mongo-template/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository interface {
	Create(*models.Product) error
	FindAll() ([]*models.Product, error)
	FindOneById(string) (*models.Product, error)
	Update(interface{}, interface{}) error
	Delete(string) error
	FindOne(interface{}) (*models.Product, error)
	ProductIsExists(interface{}) bool
}


type ProductRepositoryImp struct {
	db *mgo.Session
	config *config.Configuration
}

func (repository *ProductRepositoryImp) collection() *mgo.Collection {
	return repository.db.DB(repository.config.DatabaseName).C("products")
}


func (repository *ProductRepositoryImp) Create(product *models.Product) error {
	return repository.collection().Insert(product)
}

func (repository *ProductRepositoryImp) FindAll() ([]*models.Product, error) {
	var products []*models.Product
	err := repository.collection().Find(bson.M{}).All(&products)
	return products, err
}

func (repository *ProductRepositoryImp) FindOneById(id string) (*models.Product, error) {
	var product models.Product
	query := bson.M{"id": bson.ObjectIdHex(id)}
	err := repository.collection().Find(query).One(&product)
	return &product, err
}

func (repository *ProductRepositoryImp) Update(query, change interface{}) error {
	return repository.collection().Update(query, change)
}

func (repository *ProductRepositoryImp) Delete(id string) error {
	return repository.collection().RemoveId(bson.ObjectIdHex(id))
}

func (repository *ProductRepositoryImp) FindOne(query interface{}) (*models.Product, error) {
	var product models.Product
	err := repository.collection().Find(query).One(&product)
	return &product, err
}

func (repository *ProductRepositoryImp) ProductIsExists(query interface{}) bool {
	_, err := repository.FindOne(query)
	if err != nil {
		return false
	}
	return true
}

func NewProduct(db *mgo.Session, c *config.Configuration) ProductRepository {
	return &ProductRepositoryImp{db: db, config: c}
}