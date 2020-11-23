aws dynamodb update-table \
  --endpoint-url http://localhost:8000 \
  --table-name STMTCore \
  --global-secondary-index-updates \
    "[{\"Delete\":{\"IndexName\": \"RecordOpenness\"}}]"