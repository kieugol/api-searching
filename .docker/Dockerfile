# Use the official Golang image as a base
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on
ENV APP_ENV=development

# Install necessary packages
RUN apk add --no-cache bash curl g++ libc-dev autoconf automake libtool make

# Create and set the working directory
RUN mkdir -p /go/src/github.com/coding-challenge/api-searching
WORKDIR /go/src/github.com/coding-challenge/api-searching

COPY go.mod ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY . .

# Install CompileDaemon for hot reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 8080
 Command to run the binary
CMD ["./api"]
