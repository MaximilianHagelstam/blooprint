FROM golang:1.22

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main cmd/main.go

FROM scratch

COPY --from=0 /bin/main /bin/main

EXPOSE 8080

CMD ["/bin/main"]