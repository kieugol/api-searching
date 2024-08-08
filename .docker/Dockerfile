# Use the official Golang image as a base
FROM golang:1.20-alpine as builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_ENV=development

# Install necessary packages
RUN apk add --no-cache bash curl g++ libc-dev autoconf automake libtool

WORKDIR /app

COPY go.mod ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

#-------------------------------------------------------
FROM alpine
ENV GO111MODULE=on
ENV APP_ENV production
ENV GOPROXY=https://proxy.golang.org
ENV APP_NAME api-searching
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/myapp /app/myapp

EXPOSE 8080
#Command to run the binary api
CMD ["/app/myapp"]
