FROM golang:1.20-alpine

LABEL version="1.0"
LABEL description="Ascii Art Web deployed via containerization"
LABEL maintainer="jesseomolo@gmail.com nyunja.jp@gmail.com mcomulosammy37@gmail.com"

WORKDIR /app

COPY . .

RUN go build -o dockerize

EXPOSE 8000

ENTRYPOINT ["./dockerize"]