FROM golang:1.24

# Install ImageMagick which is required to convert PDF to images
RUN apt-get update && apt-get install -y build-essential imagemagick
RUN echo -e "<policymap>\n  <policys domain=\"coder\" rights=\"read|write\" pattern=\"PDF\" />\n</policymap>" > /etc/ImageMagick-6/policy.xml

# Copy all files into /app folder
WORKDIR /app
COPY . .

# Download all dependencies and build project
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go

# Run the application
CMD ["./main"]