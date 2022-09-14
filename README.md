# Cross Origin enabling Reverse Proxy

The Cross Origin enabling Reverse Proxy is a small, fast reverse proxy server, which
can enable CORS for services which do not enable it themselves, or remove unwanted
or unneeded extra headers from responses.

## Motivation

This was built for a search engine where the intention was to lead as much traffic as
possible back to the crawled sites by showing the crawled page immediately after the
search result.

In order to allow this, CORS headers must be set, and frame options headers must be
removed. This is a reverse proxy made especially for that purpose.

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

## Attribution

There are many implementations of reverse proxies in Go, and reading them helped
in the development of this product. In particular, [this](https://dev.to/b0r/implement-reverse-proxy-in-gogolang-2cp4)
post by b0r helped a lot and formed the baseline for the development.
