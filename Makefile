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
helm_install:
	helm install -f k8s/bff.yaml bff ./k8s/microservice
	helm install -f k8s/query-service.yaml query-service ./k8s/microservice
	helm install -f k8s/command-service.yaml command-service ./k8s/microservice

helm_uninstall:
	helm uninstall bff
	helm uninstall query-service
	helm uninstall command-service

.PHONY: build push_images build_push minikube deploy
