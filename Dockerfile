FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go mod download
RUN cd cmd/HelmCICDGenerator && \
    CGO_ENABLE=0 GOOS=linux go build -o ../../HelmCICDGenerator && \
    cd ../..

FROM redhat/ubi8:latest
COPY --from=build ./helm .
COPY --from=build ./pipelines .
COPY --from=build ./HelmCICDGenerator .
EXPOSE 8081
ENTRYPOINT [ "./HelmCICDGenerator" ]