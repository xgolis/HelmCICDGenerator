FROM golang:latest
WORKDIR /tmp
COPY . .
RUN go build -o /tmp
CMD ["/tmp"]