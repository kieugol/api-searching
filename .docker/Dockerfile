# Use the official Golang image as a base
FROM golang:1.20-alpine as builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_ENV=development

# Install necessary packages
RUN apk add --no-cache bash curl g++ libc-dev autoconf automake libtool make

# Create and set the working directory
RUN mkdir -p /go/src/github.com/coding-challenge/api-searching

WORKDIR /go/src/github.com/coding-challenge/api-searching

# Copy the rest of the application code
COPY . .
RUN ls -la

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

RUN go build -o build/api ./main.go

#-------------------------------------------------------
FROM alpine
ENV GO111MODULE=on
ENV APP_ENV production
ENV GOPROXY=https://proxy.golang.org
ENV APP_NAME api-searching
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore

RUN mkdir /usr/local/lib/searching
WORKDIR /usr/local/lib/searching

COPY --from=builder /go/src/github.com/coding-challenge/api-searching/config /usr/local/lib/searching/config
COPY --from=builder /go/src/github.com/coding-challenge/api-searching/build/api /usr/local/bin/api

EXPOSE 8080
#Command to run the binary api
ENTRYPOINT ["/usr/local/bin/api"]
