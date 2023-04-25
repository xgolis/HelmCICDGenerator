FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN cd cmd/HelmCICDGenerator && \
    CGO_ENABLE=0 GOOS=linux go build -o ../../HelmCICDGenerator && \
    cd ../..
EXPOSE 8081
ENTRYPOINT [ "./HelmCICDGenerator" ]