FROM golang:1.13.3-alpine3.10 as build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:3.10
RUN apk add --no-cache tzdata
ENV TZ Asia/Tokyo
COPY --from=build /build/main /usr/local/bin/main
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/main"]