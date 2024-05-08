FROM golang:latest
ENV API_KEY=$KEY
WORKDIR /go/src/app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
