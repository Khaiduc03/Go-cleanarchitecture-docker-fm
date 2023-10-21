
# Get Go image from DockerHub.
FROM golang:1.21-alpine AS api

# Set working directory.
WORKDIR /app

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /app

# Copy our compiled executable from the last stage.
COPY --from=api /app/server .
COPY --from=api /app/.env .
COPY --from=api /app/firebase.json .

# Run application and expose port 9999.
EXPOSE 7000
CMD ["./server"]