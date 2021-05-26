package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)


var (
	privateKey *rsa.PrivateKey
)

// Configuration contains static info required to run the apps
type Configuration struct {
	Address      string `env:"ADDRESS" envDefault:":4000"`
	MongoUrl     string `env:"MONGO_URL,required"`
	DatabaseName string `env:"DATABASE_NAME,required"`
	DbProduct    string `env:"DB_PRODUCT" envDefault:"products"`
	USERNAME     string `env:"USERNAME"`
	PASSWORD     string `env:"PASSWORD"`
}

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Println("No .env file could be found %q\n", files)
	}

	conf := Configuration{}

	err = env.Parse(&conf)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &conf
}

func GetPrivateKeyRsa() *rsa.PrivateKey {
	return privateKey
}

func InitPrivateKeyRsa() {
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
}