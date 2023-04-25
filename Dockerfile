FROM golang:latest
COPY . .
RUN CGO_ENABLE=0 go install \
    -v --work \
    ./cmd/...
RUN env
COPY ~/go/HelmCICDGenerator .
ENTRYPOINT [ "/usr/local/bin/HelmCICDGenerator" ]