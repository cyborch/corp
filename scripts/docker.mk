.PHONY: docker-build docker-push docker-clean docker-pushrm

docker-build:
	docker build -f build/Dockerfile -t $(IMAGE_NAME) $(ROOT_DIR)

docker-push:
	docker push $(IMAGE_NAME)

docker-clean:
	docker rmi $(IMAGE_NAME) || true

# docker-pushrm should be installed, see:
# https://github.com/christian-korneck/docker-pushrm
docker-pushrm:
	docker pushrm $(IMAGE_NAME) -f $(ROOT_DIR)docs/docker_hub.md -s "Cross Origin enabling Reverse Proxy"
