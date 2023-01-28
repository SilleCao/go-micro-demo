FROM golang:1.19.4-buster

# RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /build/app .

# COPY /build/app /usr/src/app

# RUN chmod -R 777 /usr/src/app

EXPOSE 7081
ENTRYPOINT ["/build/app", "run", "main.go"]