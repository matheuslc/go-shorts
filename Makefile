all: clean get-deps build test

test:
	mkdir -p bin
	go test -v -short -coverprofile=bin/cov.out ./...
	go tool cover -func=bin/cov.out

clean:
	rm -rf ./bin

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube