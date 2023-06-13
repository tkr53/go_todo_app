FROM golang:1.182-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-s -s" -o app

FROM debian:bulleye-slim as deploy

RUN apt-get update

COPY --from=deploy-bulider /app/app .

CMD ["./app"]

FROM golang:1.18.2 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]