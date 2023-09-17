FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o booking

EXPOSE 8080

CMD ["./booking"]