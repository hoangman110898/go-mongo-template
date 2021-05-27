package services

import (
	"context"
	"go-mongo-template/src/models"
	"gopkg.in/mgo.v2"
)

type UserService interface {
	Create(context.Context, *models.User) error
	Update(context.Context, *models.User) error
	Detail(ctx context.Context, user *models.User) error
}

type UserServiceImp struct {
	db *mgo.Session
}