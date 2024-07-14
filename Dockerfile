## Stage: Build
FROM golang:1.22.4-alpine AS build

WORKDIR /app

COPY ./journey .

RUN go mod download && go mod verify

RUN go build -o ./bin/journey ./cmd/journey/journey.go

## Stage: RUN
FROM scratch

WORKDIR /app

COPY --from=build /app/bin/journey/. .

EXPOSE 8080

ENTRYPOINT [ "./journey" ]