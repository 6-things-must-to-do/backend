.PHONY: clean build deploy

dev: clean 
	go build -o bin/dev ./cmd/dev/main.go
	sudo service stmt restart

dynamodb:
	docker container run -p 8000:8000 -d --name stmtcore --rm amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath .;
	./scripts/localDbInit.sh && \
 	./scripts/addInvertedGSI.sh; \
 	./scripts/addAppIDGSI.sh

local:
	go build -o bin/local ./cmd/local/main.go
	./bin/local

hot:
	reflex -s -r '\.go$$' make local

build: clean
	env GOOS=linux go build -ldflags="-s -w" -o bin/api cmd/api/main.go

clean:
	rm -rf ./bin

deploy: build
	sls deploy --verbose