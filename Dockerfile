FROM golang:1.20-alpine
WORKDIR /Pop
COPY . .
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080