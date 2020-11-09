.PHONY: build clean deploy

dev:
	go build -o bin/local ./cmd/local/main.go;
	./bin/local

dynamodb:
	docker container run -p 8000:8000 -d --name stmtcore --rm amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath .;
	./scripts/localDbInit.sh && ./scripts/addInvertedGSI.sh && ./scripts/addAppIDGSI.sh

local:
	reflex -s -r '\.go$$' make dev

build: clean
	env GOOS=linux go build -ldflags="-s -w" -o bin/api cmd/api/main.go

clean:
	rm -rf ./bin

deploy: build
	sls deploy --verbose