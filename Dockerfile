FROM golang:1.24


# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Install TypeScript globally
RUN npm install -g typescript

# Copy all files into /app folder
WORKDIR /app
COPY . .

# Download all dependencies and build project
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go

WORKDIR /app/pkg/dependencies/runtimes/typescript
RUN npm install
RUN npm install node-webcrypto-ossl
WORKDIR /app

# Run the application
CMD ["./main"]