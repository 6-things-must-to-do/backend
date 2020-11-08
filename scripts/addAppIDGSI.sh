aws dynamodb update-table \
  --endpoint-url http://localhost:8000 \
  --table-name STMTCore \
  --attribute-definitions AttributeName=SK,AttributeType=S AttributeName=AppID,AttributeType=S \
  --global-secondary-index-updates \
    "[{\"Create\":{\"IndexName\": \"AppID\",\"KeySchema\":[{\"AttributeName\":\"SK\",\"KeyType\":\"HASH\"}, {\"AttributeName\":\"AppID\",\"KeyType\":\"RANGE\"}], \
    \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 5, \"WriteCapacityUnits\": 5},\"Projection\":{\"ProjectionType\":\"ALL\"}}}]"