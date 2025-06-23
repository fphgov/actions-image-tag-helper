ARG REGISTRY=harbor.budapest.hu/docker-hub # docker.io/library

FROM ${REGISTRY}/golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["/usr/local/bin/app"]
