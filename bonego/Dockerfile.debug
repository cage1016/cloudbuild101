# build stage
FROM golang:alpine AS build-env

RUN apk update && \
    apk add --no-cache git openssh ca-certificates curl bash make build-base postgresql-client musl-dev && \
    curl https://glide.sh/get | sh && \
    go get -u github.com/derekparker/delve/cmd/dlv && \
    rm -rf /var/cache/apk/* && rm -rf /var/lib/apt/lists/*

ADD . /src
WORKDIR /src
RUN glide i && CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o app 

# final stage
FROM alpine:3.8
# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=build-env /src/app /app/
COPY --from=builder /go/bin/dlv /

EXPOSE 8080

CMD ["/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "./app"]