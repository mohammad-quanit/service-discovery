FROM golang:1.20.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

WORKDIR /service-discovery

# ENV SOURCES /service-discovery
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN go install github.com/hashicorp/consul/api@latest

RUN cd /service-discovery/client/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR /service-discovery/client/
CMD /service-discovery/client/client