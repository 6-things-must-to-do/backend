package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"github.com/labstack/gommon/log"
)

type DB struct {
	DynamoDB  *dynamo.DB
	CoreTable dynamo.Table
	SDK       *dynamodb.DynamoDB
}

var db *DB

func GetDB() *DB {
	if db != nil {
		return db
	}

	db = new(DB)

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-2"), Endpoint: aws.String("http://127.0.0.1:8000")})
	if err != nil {
		log.Error(err)
		panic(err)
	}

	sdk := dynamodb.New(awsSession)

	db.DynamoDB = dynamo.New(awsSession)
	db.CoreTable = db.DynamoDB.Table("STMTCore")
	db.SDK = sdk
	return db
}
