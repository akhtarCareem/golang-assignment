FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/userservice ./cmd/userservice/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bookingservice ./cmd/bookingservice/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/rideservice ./cmd/rideservice/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/userservice /app/bookingservice /app/rideservice .
ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_NAME=careem
CMD ["/app/userservice"]