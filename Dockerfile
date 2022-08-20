FROM golang:1.19-alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY . /app
WORKDIR /app

RUN go build -o /rental-service cmd/main.go

CMD /wait && /rental-service