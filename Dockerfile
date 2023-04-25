FROM golang:latest
COPY . .
RUN CGO_ENABLE=0 go install \
    -v --work \
    ./cmd/...
COPY /gobin/HelmCICDGenerator /usr/local/bin/HelmCICDGenerator
ENTRYPOINT [ "/usr/local/bin/HelmCICDGenerator" ]