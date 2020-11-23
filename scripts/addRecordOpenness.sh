aws dynamodb update-table \
  --endpoint-url http://localhost:8000 \
  --table-name STMTCore \
  --attribute-definitions AttributeName=SK,AttributeType=S AttributeName=RecordOpenness,AttributeType=N \
  --global-secondary-index-updates \
    "[{\"Create\":{\"IndexName\": \"RecordOpenness\",\"KeySchema\":[{\"AttributeName\":\"SK\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"RecordOpenness\",\"KeyType\":\"RANGE\"}], \
    \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 5, \"WriteCapacityUnits\": 5},\"Projection\":{\"ProjectionType\":\"ALL\"}}}]"