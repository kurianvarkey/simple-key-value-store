FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd

FROM scratch 

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]