aws dynamodb create-table \
--endpoint-url http://localhost:8000 \
--table-name STMTCore \
--attribute-definitions AttributeName=PK,AttributeType=S AttributeName=SK,AttributeType=S \
--key-schema AttributeName=PK,KeyType=HASH AttributeName=SK,KeyType=RANGE \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
--local-secondary-indexes  "[{\"IndexName\": \"Score\",
	   \"KeySchema\":[{\"AttributeName\":\"PK\",\"KeyType\":\"HASH\"},
                      {\"AttributeName\":\"SK\",\"KeyType\":\"RANGE\"}],
	   \"Projection\":{\"ProjectionType\":\"KEYS_ONLY\"}}]"
