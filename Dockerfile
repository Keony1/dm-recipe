FROM golang:1.15-buster

WORKDIR /usr/src

COPY . .

EXPOSE 8080

WORKDIR /usr/src/bin

CMD ["./dm-recipe"]