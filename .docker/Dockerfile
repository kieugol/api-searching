# Use the official Golang image as a base
FROM golang:1.20-alpine as builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_ENV=development

# Install necessary packages
RUN apk add --no-cache bash curl g++ libc-dev autoconf automake libtool

WORKDIR /app

COPY ../go.mod ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY ../. .

RUN go build -ldflags '-s -w' -a -installsuffix cgo -o build/server ./main.go

#-------------------------------------------------------
FROM alpine

RUN apk add ca-certificates tzdata
RUN mkdir /usr/local/lib/api
WORKDIR /usr/local/lib/api

ENV GO111MODULE=on
ENV APP_ENV production
ENV GOPROXY=https://proxy.golang.org
ENV APP_NAME api-searching

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/config /usr/local/lib/api/config
COPY --from=builder /app/build/server /usr/local/bin/server

EXPOSE 8080
#Command to run the binary api
CMD ["/usr/local/bin/server"]
