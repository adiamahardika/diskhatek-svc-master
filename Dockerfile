FROM golang:1.21.6 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-binary
# RUN go build -o /go-binary
# RUN go build -o /go-binary main.go

# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /
COPY --from=build-stage /go-binary /go-binary
COPY --from=build-stage /app/cmd/.env /.env


EXPOSE 3002
# USER nonroot:nonroot

# CMD ["/go-binary"]
ENTRYPOINT ["/go-binary"]
