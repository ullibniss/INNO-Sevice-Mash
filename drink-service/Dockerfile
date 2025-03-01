FROM golang:1.22.4-bullseye as builder

RUN mkdir /drink-service

WORKDIR /drink-service

COPY ./ /drink-service

RUN go mod tidy

RUN go build -o build/drink-service

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /drink-service/build/drink-service ./

CMD ["./drink-service"]