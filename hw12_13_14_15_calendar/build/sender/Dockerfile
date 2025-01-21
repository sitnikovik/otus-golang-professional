# Stage 1: Build the Go application
FROM golang:1.22 as build

ENV BIN_FILE /opt/sender/sender-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Cache dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY . ${CODE_DIR}

# Build the Go application
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/sender/*

# Stage 2: Create a minimal image with the built binary
FROM alpine:3.9

ENV BIN_FILE "/opt/sender/sender-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

# Copy the configuration file
ENV CONFIG_FILE /etc/sender/.env
COPY .env ${CONFIG_FILE}

# Command to run the application
CMD ${BIN_FILE} -config ${CONFIG_FILE}