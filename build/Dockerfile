FROM golang:1.19-alpine as build

RUN apk add --update make

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./
RUN make corp

FROM golang:1.19-alpine as final

WORKDIR /app
RUN mkdir -p bin 
COPY --from=build /app/bin/corp /app/bin/corp
COPY ./configs ./configs

CMD /app/bin/corp
