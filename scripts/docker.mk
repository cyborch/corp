.PHONY: docker-build docker-push docker-clean

docker-build:
	docker build -f build/Dockerfile -t $(IMAGE_NAME) $(ROOT_DIR)

docker-push:
	docker push $(IMAGE_NAME)

docker-clean:
	docker rmi $(IMAGE_NAME) || true
