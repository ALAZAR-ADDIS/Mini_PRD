#!/bin/bash

# Blog Platform Setup Script
echo "🚀 Setting up Blog Platform..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.24.4 or higher."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
if [[ $(echo "$GO_VERSION 1.24.4" | tr " " "\n" | sort -V | head -n 1) != "1.24.4" ]]; then
    echo "❌ Go version $GO_VERSION is too old. Please install Go 1.24.4 or higher."
    exit 1
fi

echo "✅ Go version $GO_VERSION is compatible"

# Download dependencies
echo "📦 Downloading dependencies..."
go mod download

# Check if .env file exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp env.example .env
    echo "⚠️  Please update .env file with your MongoDB Atlas credentials"
    echo "📖 See Database/mongodb_atlas_setup.md for detailed instructions"
else
    echo "✅ .env file already exists"
fi

# Check if MongoDB URI is configured
if grep -q "your_username" .env; then
    echo "⚠️  Please update MONGODB_URI in .env file with your MongoDB Atlas credentials"
fi

# Check if JWT secrets are configured
if grep -q "your_super_secret" .env; then
    echo "⚠️  Please update JWT_SECRET and JWT_REFRESH_SECRET in .env file"
fi

echo ""
echo "🎉 Setup complete!"
echo ""
echo "Next steps:"
echo "1. Update .env file with your MongoDB Atlas credentials"
echo "2. Follow Database/mongodb_atlas_setup.md to configure MongoDB Atlas"
echo "3. Run 'go run main.go' to start the server"
echo ""
echo "📚 Documentation:"
echo "- Database setup: Database/mongodb_atlas_setup.md"
echo "- API documentation: README.md"
echo "- Database design: Database/database_design.md" 