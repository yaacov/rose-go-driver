# Build stage
FROM docker.io/library/golang:1.19 AS build-env

# Set up the working directory
WORKDIR /go/src/app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s' -o rose-go-driver ./cmd/

# Final stage
FROM scratch

# Copy the compiled binary from the build stage
COPY --from=build-env /go/src/app/rose-go-driver /app/

ENTRYPOINT ["/app/rose-go-driver", "-listen", "0.0.0.0"]
CMD ["-port", "8081"]
