FROM golang:1.22.4-bullseye as builder

RUN mkdir /syrup-service

WORKDIR /syrup-service

COPY ./ /syrup-service

RUN go mod tidy

RUN go build -o build/syrup-service

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /syrup-service/build/syrup-service ./

CMD ["./syrup-service"]