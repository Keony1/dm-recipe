FROM golang:1.15-buster

RUN mkdir -p /usr/src
WORKDIR /usr/src

COPY . ./

EXPOSE 8080

CMD ["/usr/src/recipe"]