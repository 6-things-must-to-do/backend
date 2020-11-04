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

func InitDB(isLocal bool) *DB {
	d := new(DB)

	d.DynamoDB = dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-2"), Endpoint: aws.String("http://localhost:8000")})
	d.CoreTable = d.DynamoDB.Table("STMTCore")
	return d
}
