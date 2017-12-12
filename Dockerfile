FROM golang:1.9

COPY *.go /go/
RUN go build -o /home/server

ENTRYPOINT ["/home/server"]

EXPOSE 8080
