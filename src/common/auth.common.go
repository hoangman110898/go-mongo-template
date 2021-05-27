package common

import "go-mongo-template/src/models"

type LoginRes struct {
	Token string `json:"token"`
	User *models.User `json:"user"`
}