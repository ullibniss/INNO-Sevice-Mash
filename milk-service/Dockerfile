FROM golang:1.22.4-bullseye as builder

RUN mkdir /milk-service

WORKDIR /milk-service

COPY ./ /milk-service

RUN go mod tidy

RUN go build -o build/milk-service

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /milk-service/build/milk-service ./

CMD ["./milk-service"]