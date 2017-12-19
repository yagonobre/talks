FROM golang:alpine
RUN mkdir /app
ADD ping.go /app
RUN go build -o /app/ping /app/ping.go
CMD ["/app/ping"]

EXPOSE 8080



