package models

import (
	"go-mongo-template/src/utility"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	Price string        `json:"price,omitempty"`
	Code  string        `json:"code,omitempty"`
	Image string        `json:"image,omitempty"`
}

func (p *Product) Validate() error {
	if e := utility.ValidateRequireAndLengthAndRegex(p.Name, true, 3, 25, "", "Name"); e != nil {
		return e
	}

	if e := utility.ValidateRequireAndLengthAndRegex(p.Code, true, 2, 20, "", "Code"); e != nil {
		return e
	}

	return nil
}
