FROM golang:alpine

WORKDIR /build

COPY . .

CMD ["go", "run", "cmd/app/main.go"]