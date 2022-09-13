# Cross Origin enabling Reverse Proxy

The Cross Origin enabling Reverse Proxy is a small, fast reverse proxy server, which
can enable CORS for services which do not enable it themselves, or remove unwanted
or unneeded extra headers from responses.

This is especially useful when accessing services which do not allow shoing their
content in iframes (i.e. by setting `X-Frame-Options`).

## Configuration

The service is configured in a `proxy.yml` file, see [configs](configs) for details.

## Building

Build script details can be found in the [scripts](scripts) folder.

Build local binary by running:

```
make all
```

## Deployment

### Docker

The `IMAGE_NAME` in the root `Makefile` specifies the docker image tag used.
Update it as appropriate before building docker images.

Build docker and push images by running:

```
make docker-build
make docker-push
```

### Helm

The `image` and `virtualHosts` should be updated appropriately in 
[values.yaml](deployments/helm/values.yaml).

Once docker images are built and pushed, as seen above, then deploy the service
by running:

```
make helm-install
```
