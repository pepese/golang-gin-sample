FROM alpine:3.8
EXPOSE 8080
ADD app app
ENTRYPOINT ["./app"]