# Multistage docker build 
# step-1 : Build the Go application

# 1. need some OS basically linux distribution ? Docker works only on linux 
# 2. Golang to be installed on docker 

FROM golang:1.26-alpine AS builder
# I want an image that contains some OS(alpine)+ Golang installed onit

RUN apk add --no-cache tzdata

WORKDIR /app 

# it creates a directory called app and also cd app

COPY go.mod go.sum ./

RUN go mod download

# copy the remaining source files

COPY . .  

# build the application as a static binary 

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/bin/server ./cmd/api


# step-2 : Run the compiled binary

# FROM alpine:3.20

# WORKDIR /app

# # Copy the binary from the builder stage into this alpine stage
# COPY --from=builder /app/bin/server ./server

# # Start the interactive shell when the container runs
# CMD ["sh"]

FROM scratch

# copy timezone info
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the static binary from the 'builder' stage
COPY --from=builder /app/bin/server /server

# (Optional) If your app makes HTTPS/TLS requests, you need SSL certificates
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the port your server listens on (e.g., 8080)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/server"]

