FROM golang:1.19.10-bullseye
RUN mkdir /cdbp
WORKDIR /cdbp
COPY . /cdbp
RUN go mod tidy
ENTRYPOINT [ "go", "run", "server.go" ]