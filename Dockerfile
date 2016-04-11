FROM golang:latest 
COPY src/ /go/src/
COPY resources/ /resources/
EXPOSE 8080

WORKDIR src
RUN go build -o /app/main github.com/josdirksen/systeminfo/main.go
WORKDIR /

CMD ["/app/main"]