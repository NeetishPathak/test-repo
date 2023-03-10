FROM golang:1.18.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY server ./server
COPY *.go ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]