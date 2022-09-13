# Service Configuration

The service is configured in a `proxy.yml` file, which can look as follows:

```
server:
  port: 8080


virtualHosts:
  - hostname: localhost:8080
    scheme: http
    origin: https://google.com
    enableCors: true
    skipHeaders:
      - X-Frame-Options
```

The `server.port` configuration specifies which port the server listens on.

Virtual hosts can be specified. Each virtual host forwards
requests to a given hostname on a given protocol (scheme) to the 
specified origin.

CORS can be enabled or disabled for a given virtual host

Also, any unwanted headers can be skipped from the origin response
