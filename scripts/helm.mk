.PHONY: helm-install helm-uninstall

helm-install:
	helm install corp $(ROOT_DIR)deployments/helm

helm-uninstall:
	helm uninstall corp
