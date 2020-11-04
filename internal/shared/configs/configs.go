package configs

import "os"

type DynamoDBConfigs struct {
	ENDPOINT string
	REGION   string
}

type Config struct {
	SECRET string
	DYNAMO DynamoDBConfigs
}

var c *Config

func GetConfig() *Config {
	if c != nil {
		return c
	}

	c = &Config{SECRET: os.Getenv("SECRET"), DYNAMO: DynamoDBConfigs{
		ENDPOINT: os.Getenv("DB_ENDPOINT"),
		REGION:   os.Getenv("DB_REGION"),
	}}
	return c
}
