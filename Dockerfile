FROM golang:1.17 as build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM gcr.io/distroless/static
ENV TZ Asia/Tokyo
COPY --from=build /build/main /usr/local/bin/main
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/main"]