FROM golang:1.24.1

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go_docker_app

CMD [ "/go_docker_app" ]

