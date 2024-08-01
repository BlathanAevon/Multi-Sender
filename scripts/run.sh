echo "Downloading dependencies..."
go mod download

echo "Launching..."
go run cmd/MultiSender/main.go