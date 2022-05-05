
docker-build: #Build image
	@echo "Build image..."
	docker build -t wander4747/hello-k8s .

docker-push: #Push Docker Hub
	@echo "Pushing image to Docker Hub..."
	docker push wander4747/hello-k8s