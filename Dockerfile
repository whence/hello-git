FROM golang:1.17.1-bullseye AS build-env

RUN apt-get update && apt-get install -y gcc libgit2-1.1 libgit2-dev

ADD ["main.go", "go.mod", "go.sum", "/app/"]
RUN cd /app && go build -tags static,system_libgit2 -o server main.go

# FROM ubuntu:20.04
# COPY --from=build-env ["/app/server", "/app/"]
# WORKDIR /app
# ENTRYPOINT ["./server"]
