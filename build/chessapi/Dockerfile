FROM golang:1.22

WORKDIR /usr/src/app

COPY ../../chessapi/go.mod ../../chessapi/go.sum ./
RUN go mod download && go mod verify

COPY ../../chessapi .
RUN go build -v -o /usr/local/bin/app .

CMD ["/usr/local/bin/app"]