package products

import (
	"context"
	"go-mongo-template/src/config"
	"go-mongo-template/src/models"
	"go-mongo-template/src/repositories"
	"go-mongo-template/src/utility"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductService interface {
	Create(context.Context, *models.Product) error
	FindAll(context.Context) ([]*models.Product, error)
	FindOneById(context.Context, string) (*models.Product, error)
	Update(context.Context, string, *models.Product) error
	Delete(context.Context, string) error
	FindOne(context.Context, *models.Product) (*models.Product, error)
}

type ProductServiceImp struct {
	db         *mgo.Session
	repository repositories.ProductRepository
	config     *config.Configuration
}

func (service *ProductServiceImp) Create(ctx context.Context, product *models.Product) error {
	return service.repository.Create(ctx, product)
}

func (service *ProductServiceImp) FindAll(ctx context.Context) ([]*models.Product, error) {
	return service.repository.FindAll(ctx)
}

func (service *ProductServiceImp) FindOneById(ctx context.Context, id string) (*models.Product, error) {
	return service.repository.FindOneById(ctx, id)
}

func (service *ProductServiceImp) Update(ctx context.Context, id string, product *models.Product) error {
	query := bson.M{"_id": id}

	customBson := &utility.CustomBson{}
	change, err := customBson.Set(product)
	if err != nil {
		return err
	}
	return service.repository.Update(ctx, query, change)
}

func (service *ProductServiceImp) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}

func (service *ProductServiceImp) FindOne(ctx context.Context, product *models.Product) (*models.Product, error) {
	customBson := &utility.CustomBson{}

	find, err := customBson.Set(product)
	if err != nil {
		return nil, err
	}

	return service.repository.FindOne(ctx, find)
}

func New(db *mgo.Session, c *config.Configuration) ProductService {
	return &ProductServiceImp{
		db:         db,
		config:     c,
		repository: repositories.New(db, c),
	}
}
