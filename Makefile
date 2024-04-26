# Define variables for image names and paths
BFF_IMAGE := bff:latest
QUERY_IMAGE := query_service:latest
COMMAND_IMAGE := command_service:latest
DOCKERHUB_REPO := keigokida/gomicroservices

build:
	docker build -t $(BFF_IMAGE) ./bff
	docker build -t $(QUERY_IMAGE) ./microservices/query_service
	docker build -t $(COMMAND_IMAGE) ./microservices/command_service
	docker tag $(BFF_IMAGE) $(DOCKERHUB_REPO):bff
	docker tag $(QUERY_IMAGE) $(DOCKERHUB_REPO):query_service
	docker tag $(COMMAND_IMAGE) $(DOCKERHUB_REPO):command_service

push_images:
	docker push $(DOCKERHUB_REPO):bff
	docker push $(DOCKERHUB_REPO):query_service
	docker push $(DOCKERHUB_REPO):command_service

build_push: build push_images

minikube:
	minikube start

# Need to start minikube before deploying services
deploy:
	helm install -f bff.yaml bff ./microservice
	helm install -f query-service.yaml query-service ./microservice
	helm install -f command-service.yaml command-service ./microservice

.PHONY: build push_images build_push minikube deploy
