.PHONY: helm-install helm-uninstall helm-publish

helm-install:
	helm install corp $(ROOT_DIR)deployments/helm

helm-uninstall:
	helm uninstall corp

helm-page:
	helm plugin install https://github.com/halkeye/helm-repo-html || true
	mkdir -p $(ROOT_DIR)docs/pages/charts
	curl --silent -f -o $(ROOT_DIR)docs/pages/charts/index.yaml https://cyborch.github.io/corp/charts/index.yaml || \
		echo "apiVersion: v1" > $(ROOT_DIR)docs/pages/charts/index.yaml
	helm package -d $(ROOT_DIR)docs/pages/charts $(ROOT_DIR)deployments/helm
	helm repo index $(ROOT_DIR)docs/pages/charts \
		--url https://cyborch.github.io/corp/charts \
		--merge $(ROOT_DIR)docs/pages/charts/index.yaml
	$(ROOT_DIR)scripts/download-charts.sh
	helm repo-html \
		-i $(ROOT_DIR)docs/pages/charts/index.yaml \
		-o $(ROOT_DIR)docs/pages/index.html \
		-t $(ROOT_DIR)docs/github_page.tpl
