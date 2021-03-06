FROM golang:1.16-alpine as builder
RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client ./client

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server  .
COPY --from=builder /app/client  .

EXPOSE 8081
CMD [ "./server" ]
