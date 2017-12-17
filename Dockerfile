FROM golang:1.9

COPY *.go /go/
RUN go build -o /home/server
COPY test.sh /
RUN chmod +x /test.sh

ENTRYPOINT ["/home/server"]

EXPOSE 8080
