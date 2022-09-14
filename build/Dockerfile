FROM golang:1.19.1-alpine3.15 as build

RUN apk add --update make

WORKDIR /app

# Download project dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build corp binary
COPY ./ ./
RUN make corp

FROM alpine:3.15 as final

WORKDIR /app
RUN mkdir -p bin 
COPY --from=build /app/bin/corp /app/bin/corp

CMD /app/bin/corp
