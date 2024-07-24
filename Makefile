# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
DOCKER_IMAGE=go-gin-app
DOCKER_CONTAINER=go-gin-app

# Build the Go project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Clean the Go project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
docker-run:
	docker run -d -p 8080:8080 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop Docker container
docker-stop:
	docker stop $(DOCKER_CONTAINER)

# Remove Docker container
docker-rm:
	docker rm $(DOCKER_CONTAINER)

# Remove Docker image
docker-rmi:
	docker rmi $(DOCKER_IMAGE)

.PHONY: build clean test docker-build docker-run docker-stop docker-rm docker-rmi
