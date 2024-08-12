# Use the official Golang image as a base
FROM golang:1.20-alpine as builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_ENV=development

# Install necessary packages
RUN apk add --no-cache bash curl g++ libc-dev autoconf automake libtool

RUN ls -la

WORKDIR /app

RUN ls -la

COPY go.* ./

# Tidy up module dependencies and create vendor directory
RUN go mod tidy -compat=1.20
RUN go mod vendor

# Copy the rest of the application code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

#-------------------------------------------------------
FROM scratch

ENV GO111MODULE=on
ENV APP_ENV production
ENV GOPROXY=https://proxy.golang.org
ENV APP_NAME api-searching
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore

RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/config ./app/config
COPY --from=builder /app/server ./app/server

EXPOSE 8080
#Command to run the binary api
CMD ["/app/server"]
