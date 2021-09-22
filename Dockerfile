FROM golang:1.17.1-bullseye AS build-env

RUN apt-get update && apt-get install -y gcc cmake pkg-config

ADD ["*.go", "go.*", "*.sh", "/app/"]

RUN ["/app/install_git2go.sh"]
RUN ["/app/build_app.sh"]

FROM ubuntu:20.04
COPY --from=build-env ["/app/server", "/app/"]
WORKDIR /app
ENTRYPOINT ["./server"]
