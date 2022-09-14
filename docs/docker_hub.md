# Quick reference

* **Maintained by:**<br>
  [Anders Borch](https://github.com/cyborch/corp)

* **Where to get help:**<br>
  [Project readme](https://github.com/cyborch/corp), [Github issues](https://github.com/cyborch/corp/issues).

# Supported tags and respective `Dockerfile` links

* [latest](https://github.com/cyborch/corp/blob/main/build/Dockerfile)

# What is corp?

The Cross Origin enabling Reverse Proxy is a small, fast reverse proxy server, which
can enable CORS for services which do not enable it themselves, or remove unwanted
or unneeded extra headers from responses.

## Motivation

This was built for a search engine where the intention was to lead as much traffic as
possible back to the crawled sites by showing the crawled page immediately after the
search result.

In order to allow this, CORS headers must be set, and frame options headers must be
removed. This is a reverse proxy made especially for that purpose.

# How to use this image

The service is configured in a `proxy.yml` file, which can look as follows:

```
server:
  port: 8080


virtualHosts:
  - hostname: proxy.example.com
    scheme: http
    origin: https://www.google.com
    enableCors: true
    skipHeaders:
      - X-Frame-Options
```

The `server.port` configuration specifies which port the server listens on.

Virtual hosts can be specified. Each virtual host forwards
requests to a given hostname on a given protocol (scheme) to the 
specified origin.

CORS can be enabled or disabled for a given virtual host

Also, any unwanted headers can be skipped from the origin response by setting
the `virtualHosts` array.

* `hostname` the hostname where the corp service is hosted.
* `scheme` the request scheme used (e.g. `http` or `https`).
* `origin` the name of the service which requests are forwarded to.
* `enableCors` sets CORS headers in the forwarded response from origin if true.
* `skipHeaders` removes any named headers in the forwarded response from origin.

Place this file in the `configs` directory, then start your container:

```
$ docker run --name my-custom-corp-container -v /host/path/configs/proxy.yml:/app/configs/proxy.yml:ro -d cyborch/corp
```

## Helm deployment

Specify a docker `image` name, aand a list of virtual hosts in `values.yaml`:

```
namespace: default
image: cyborch/corp

virtualHosts:
  - hostname: proxy.example.com
    scheme: http
    origin: https://www.google.com
    enableCors: true
    skipHeaders:
      - X-Frame-Options
```

Place this file in the `deployments/helm` directory, then deploy your chart:

```
$ helm install corp ./deployments/helm
```

# License

This software is covered by the MIT license.

View [license information](https://github.com/cyborch/corp/blob/main/LICENSE) for the software contained in this image.
