.PHONY: build clean deploy

devbuild:
	go build -o bin/local ./cmd/local/main.go;
	./bin/local

dynamodb:
	docker container run -p 8000:8000 -d --name stmtcore --rm -v /User/changhoi/dev/6-things-must-to-do amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath .

local:
	reflex -s -r '\.go$$' make devbuild

build: clean
	env GOOS=linux go build -ldflags="-s -w" -o bin/api cmd/api/main.go

clean:
	rm -rf ./bin

deploy: build
	sls deploy --verbose