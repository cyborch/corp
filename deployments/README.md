# Deployments

This folder contains kubernetes deployment infrastructure information.

Specify a docker `image` name, aand a list of virtual hosts
in [values.yaml](helm/values.yaml).

See `configs` for details.

To build, push, and deploy CORP, run the following in the root folder:

```
make docker-build docker-push helm-install
```
