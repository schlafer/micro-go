# Use the updated Go version 1.22.7 with Alpine
FROM golang:1.22.7-alpine3.20 AS build

# Install necessary tools
RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory
WORKDIR /go/src/github.com/schlafer/micro-go

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY vendor vendor
COPY account account
COPY catalog catalog
COPY order order

# Build the application targeting the order module
RUN go build -mod=vendor -o /go/bin/app ./order/cmd/order

# Use a minimal Alpine image for the runtime environment
FROM alpine:3.20

# Set the working directory
WORKDIR /usr/bin

# Copy the built application from the build stage
COPY --from=build /go/bin .

# Expose the application's port
EXPOSE 8080

# Set the command to run the application
CMD ["app"]
