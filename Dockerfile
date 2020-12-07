FROM golang:1.15-buster

WORKDIR /usr/src

COPY . .

EXPOSE 8080

RUN mkdir -p bin

WORKDIR /usr/src/bin

CMD ["./dm-recipe"]