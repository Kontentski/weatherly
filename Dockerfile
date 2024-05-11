FROM golang:latest
ENV API_KEY=$KEY
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /weatherly
EXPOSE 8080
CMD ["/weatherly"]
