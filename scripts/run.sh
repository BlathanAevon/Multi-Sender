echo "Downloading dependencies..."
go mod download

echo "Launching..."
go run cmd/tabi/main.go