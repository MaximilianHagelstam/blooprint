FROM golang:1.22-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /bin/main cmd/main.go

FROM scratch

COPY --from=0 /bin/main /bin/main

EXPOSE 8080

CMD ["/bin/main"]