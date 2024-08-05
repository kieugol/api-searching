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

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY . .

# Install CompileDaemon for hot reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set the entrypoint to use CompileDaemon for hot reloading
ENTRYPOINT ["CompileDaemon", "-build=go build -o build/api ./main.go", "-command=./build/api"]
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

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY . .

RUN go build -o build/api ./main.go

EXPOSE 8088
# Set the entrypoint to use CompileDaemon for hot reloading
CMD ["./api"]
