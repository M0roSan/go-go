# base image
FROM golang:1.13 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# build go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

#Copy the pre-build binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]