# Careem Microservices (Go, gRPC, Postgres)

## Overview

This is a simplified microservices-based application for a Careem-like system using:
- Go
- gRPC with Protobuf
- PostgreSQL database
- Logging with logrus
- Prometheus metrics

## Services

- UserService: Manages users
- BookingService: Creates and retrieves bookings
- RidesService: Updates rides

## How to Run

1. Install dependencies:
   ```bash
   go mod tidy