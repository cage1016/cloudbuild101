# build stage
FROM golang:alpine AS build-env

RUN apk update && \
    apk add --no-cache git openssh ca-certificates curl bash make build-base postgresql-client && \
    curl https://glide.sh/get | sh && \
    rm -rf /var/cache/apk/* && rm -rf /var/lib/apt/lists/*

ADD . /src
WORKDIR /src
RUN glide i && CGO_ENABLED=0 GOOS=linux go build -o app 

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-env /src/app /app/

EXPOSE 8080

CMD ["./app"]