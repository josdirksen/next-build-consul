FROM golang:latest 

COPY src/ /go/src/
COPY resources/ /resources/
COPY scripts/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

EXPOSE 8080

WORKDIR src
RUN go build -o /app/main github.com/josdirksen/nbdemo/main.go
WORKDIR /

CMD ["/entrypoint.sh"]