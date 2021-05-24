package repositories

import (
	"context"
	"go-mongo-template/src/config"
	"go-mongo-template/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository interface {
	Create(context.Context, *models.Product) error
	FindAll(context.Context) ([]*models.Product, error)
	FindOneById(context.Context, string) (*models.Product, error)
	Update(context.Context, interface{}, interface{}) error
	Delete(context.Context, string) error
	FindOne(context.Context, interface{}) (*models.Product, error)
	ProductIsExists(context.Context, interface{}) bool
}


type ProductRepositoryImp struct {
	db *mgo.Session
	config *config.Configuration
}

func (repository *ProductRepositoryImp) collection() *mgo.Collection {
	return repository.db.DB(repository.config.DatabaseName).C("products")
}


func (repository *ProductRepositoryImp) Create(_ context.Context, product *models.Product) error {
	return repository.collection().Insert(product)
}

func (repository *ProductRepositoryImp) FindAll(_ context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := repository.collection().Find(bson.M{}).All(&products)
	return products, err
}

func (repository *ProductRepositoryImp) FindOneById(_ context.Context, id string) (*models.Product, error) {
	var product models.Product
	query := bson.M{"id": bson.ObjectIdHex(id)}
	err := repository.collection().Find(query).One(&product)
	return &product, err
}

func (repository *ProductRepositoryImp) Update(_ context.Context, query, change interface{}) error {
	return repository.collection().Update(query, change)
}

func (repository *ProductRepositoryImp) Delete(_ context.Context, id string) error {
	return repository.collection().RemoveId(bson.ObjectIdHex(id))
}

func (repository *ProductRepositoryImp) FindOne(_ context.Context, query interface{}) (*models.Product, error) {
	var product models.Product
	err := repository.collection().Find(query).One(&product)
	return &product, err
}

func (repository *ProductRepositoryImp) ProductIsExists(ctx context.Context, query interface{}) bool {
	_, err := repository.FindOne(ctx, query)
	if err != nil {
		return false
	}
	return true
}

func New(db *mgo.Session, c *config.Configuration) ProductRepository {
	return &ProductRepositoryImp{db: db, config: c}
}