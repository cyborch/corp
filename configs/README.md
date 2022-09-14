# Service Configuration

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
