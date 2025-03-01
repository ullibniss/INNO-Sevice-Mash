FROM golang:1.22.4-bullseye as builder

RUN mkdir /coffee-service

WORKDIR /coffee-service

COPY ./ /coffee-service

RUN go mod tidy

RUN cd cmd && \ 
    go build -o ../build/coffee-service

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /coffee-service/build/coffee-service ./

CMD ["./coffee-service"]