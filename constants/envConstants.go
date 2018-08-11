package constants

import (
	en "github.com/caarlos0/env"
	"os"
)

//Env contains all the environment variables
var Env envVars

func init() {
	err := en.Parse(&Env)
	if err != nil {
		os.Exit(1)
	}
}

type envVars struct {

	//server
	Port string `env:"PORT" envDefault:"8489"`

	//aws configs
	AwsAccessKeyID     string `env:"AWS_ACCESS_KEY_ID" envDefault:""`
	AwsSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" envDefault:""`
	AwsRegion          string `env:"AWS_REGION" envDefault:"us-east-1"`
	AwsBucket          string `env:"AWS_BUCKET" envDefault:"ts-engineering-test"`

	//mongo configs
	MongoURL        string `env:"MONGO_URL" envDefault:""`
	MongoDBName     string `env:"MONGO_DB_NAME" envDefault:"dblabs"`
	MongoCollection string `env:"MONGO_COLLECTION" envDefault:"assets"`
}
