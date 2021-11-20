FROM golang:1.16-alpine
WORKDIR /app
RUN apk add --no-cache git
COPY go.sum ./
COPY go.mod ./
RUN go mod download
RUN go mod tidy
COPY *.go ./
COPY casino/*.go ./casino

RUN go build -o /casinowar
EXPOSE 8081
CMD [ "/casinowar", "8081" ]
