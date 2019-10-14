FROM golang:1.13-alpine

ARG GOPROXY

WORKDIR /build

ADD . /build

RUN GOPROXY=$GOPROXY CGO_ENABLED=0 go build -o gobundle ./cmd/gobundle

FROM scratch

COPY --from=0 /build/gobundle /gobundle

ENTRYPOINT [ "/gobundle" ]
