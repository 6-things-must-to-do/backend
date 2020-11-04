package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DB struct {
	DynamoDB  *dynamo.DB
	CoreTable dynamo.Table
}

var db *DB

func GetDB() *DB {
	if db != nil {
		return db
	}

	db = new(DB)

	db.DynamoDB = dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-2"), Endpoint: aws.String("http://localhost:8000")})
	db.CoreTable = db.DynamoDB.Table("STMTCore")
	return db
}
