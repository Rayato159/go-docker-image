<h1>Golang Basic Dockerfile Example 🐳</h1>

<h2>For smallest image</h2>
***This image not avaliable for some function such as create file.

```docker
FROM golang:1.20-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/app /bin

EXPOSE 1323

ENTRYPOINT [ "/bin/app" ]
```

<h2>For basic debian image</h2>

```docker
FROM golang:1.20-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /bin/app

FROM debian:buster-slim

COPY --from=build /bin/app /bin

EXPOSE 1323

CMD [ "/bin/app" ]
```