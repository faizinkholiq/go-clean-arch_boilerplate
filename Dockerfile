FROM golang:1.23-alpine

RUN apk add --no-cache git curl build-base

# Set the working directory inside the container
WORKDIR /app

ENV GOFLAGS="-buildvcs=false"

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Expose the application port
EXPOSE 8080

# Use Air to run the application in development mode
CMD ["air", "-c", ".air.toml", "-d"]