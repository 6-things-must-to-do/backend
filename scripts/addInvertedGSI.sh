aws dynamodb update-table \
  --endpoint-url http://localhost:8000 \
  --table-name STMTCore \
  --attribute-definitions AttributeName=SK,AttributeType=S AttributeName=PK,AttributeType=S \
  --global-secondary-index-updates \
    "[{\"Create\":{\"IndexName\": \"Inverted\",\"KeySchema\":[{\"AttributeName\":\"SK\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"PK\",\"KeyType\":\"RANGE\"}], \
    \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 5, \"WriteCapacityUnits\": 5},\"Projection\":{\"ProjectionType\":\"ALL\"}}}]"