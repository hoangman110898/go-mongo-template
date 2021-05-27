package services

import (
	"github.com/form3tech-oss/jwt-go"
	"go-mongo-template/src/common"
	"go-mongo-template/src/config"
	"go-mongo-template/src/models"
	"go-mongo-template/src/repositories"
	"go-mongo-template/src/utility"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type UserService interface {
	Create(*models.User) error
	Update(string, *models.User) error
	Detail(string) error
	FindByEmail(string) (*models.User, error)
	Login(string, string) (*common.LoginRes, error)
}

type UserServiceImp struct {
	db *mgo.Session
	repository repositories.UserRepository
	config *config.Configuration
}

func (service *UserServiceImp) Login(email string, password string) (*common.LoginRes, error) {
	userDb, _ := service.FindByEmail(email)
	if  userDb == nil {
		return nil, nil
	}
	err := userDb.ComparePassword(password)
	if err != nil {
		return nil, err
	}
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userDb.ID
	claims["name"] = userDb.Name
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	if userDb.Role == "ADMIN" {
		claims["admin"] = false
	} else {
		claims["admin"] = true
	}
	claims["admin"] = false

	t, err := token.SignedString(privateKey)

	if err != nil {
		log.Printf("token.SignedString : %v", err)
		return nil, err
	}

	result := common.LoginRes{Token: t, User: userDb}
	return &result, nil
}

func (service *UserServiceImp) FindByEmail(email string) (*models.User, error) {
	return service.repository.FindByEmail(email)
}

func (service *UserServiceImp) Create(user *models.User) error {
	return service.repository.Create(user)
}

func (service *UserServiceImp) Update(id string, user *models.User) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	customBson := &utility.CustomBson{}
	change, err := customBson.Set(user)

	if err != nil {
		return nil
	}
	return service.repository.Update(query, change)
}

func (service *UserServiceImp) Detail(id string) error {
	return service.repository.Detail(id)
}

func NewUser(db *mgo.Session, c *config.Configuration) UserService {
	return &UserServiceImp{
		db: db,
		config: c,
		repository: repositories.NewUser(db, c),
	}
}