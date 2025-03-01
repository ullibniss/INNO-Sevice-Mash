FROM golang:1.22.4-bullseye as builder

ARG GO111MODULE=on

ARG GOPROXY=https://goproxy.cn

RUN mkdir /inno-lb-sidecar

WORKDIR /inno-lb-sidecar

COPY ./ /inno-lb-sidecar

RUN go mod tidy

RUN cd ./cmd && \
    go build -o ../build/inno-lb-sidecar

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /inno-lb-sidecar/build/inno-lb-sidecar ./

COPY config.yml ./

CMD ["./inno-lb-sidecar"]