AppName := "FilesGo"

# Print this help message
@help:
  just --list

# Run the cmd/FilesGo application
[no-exit-message]
@run:
  echo "Running {{AppName}}..."
  go run ./cmd/{{AppName}}

# Only installs the imports using go get
@install:
  echo "Installing dependencies..."
  go mod tidy

# Format all .go files, and tidy and vendor module dependencies
@tidy:
  echo "Formatting .go files..."
  go fmt ./...
  just install
  echo "Verifying dependencies..."
  go mod verify
