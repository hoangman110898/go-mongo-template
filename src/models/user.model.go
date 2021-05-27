package models

import (
	"github.com/google/uuid"
	"go-mongo-template/src/utility"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Salt     string        `json:"salt"`
	Role     string        `json:"role,omitempty"`
}

func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordBytes := []byte(u.Password + salt)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash[:])
	u.Salt = salt
	u.Role = utility.UserRole
	return nil
}

func (u *User) Validate() error {
	if e := utility.ValidateRequireAndLengthAndRegex(u.Name, true, 3, 25, "", "Name"); e != nil {
		return e
	}

	if e := utility.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "Email"); e != nil {
		return e
	}

	if e := utility.ValidateRequireAndLengthAndRegex(u.Password, true, 8, 25, "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$", "Password"); e != nil {
		return e
	}

	return nil
}
