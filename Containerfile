FROM golang:1.25.4 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/smtn

FROM golang:1.25.4 AS release

WORKDIR /app
COPY --from=build /app/bin/smtn /app/bin/smtn

ENTRYPOINT ["/app/bin/smtn"]
