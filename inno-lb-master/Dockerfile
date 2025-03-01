FROM golang:1.22.4-bullseye as builder

RUN mkdir /inno-lb-master

WORKDIR /inno-lb-master

COPY ./ /inno-lb-master

RUN go mod tidy

RUN cd ./cmd && \
    go build -o ../build/inno-lb-master

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /inno-lb-master/build/inno-lb-master ./

COPY config.yml ./

CMD ["./inno-lb-master"]