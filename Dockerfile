FROM golang:1.20-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/app /bin

EXPOSE 1323

ENTRYPOINT [ "/bin/app" ]