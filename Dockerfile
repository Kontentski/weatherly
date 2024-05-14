FROM golang:latest AS builder
ENV KEY=$KEY
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags "-extldflags '-static'" -o /bin/weatherly

FROM alpine:latest
COPY --from=builder /bin/weatherly /bin/weatherly
COPY . .

EXPOSE 8080
CMD ["/bin/weatherly"]
