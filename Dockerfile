FROM golang:1.21.3

WORKDIR /go/src

CMD ["tail", "-f", "/dev/null"]
