#!/bin/bash

echo "🔍 Debug Build Process..."

# Check Go version
echo "📋 Go version:"
go version

# Check current directory and files
echo "📁 Current directory:"
pwd
ls -la

# Check go.mod
echo "📄 go.mod content:"
cat go.mod

# Clean and build
echo "🧹 Cleaning..."
go clean
rm -f main

echo "🔨 Building..."
go build -v -o main .

# Check if binary was created
if [ -f "main" ]; then
    echo "✅ Binary created successfully!"
    echo "📊 Binary info:"
    ls -la main
    file main
    
    # Test if binary is executable
    echo "🧪 Testing binary execution..."
    chmod +x main
    timeout 5s ./main --help 2>&1 || echo "Binary test completed (timeout expected)"
else
    echo "❌ Binary was not created!"
    echo "🔍 Checking for errors..."
    go build -v -o main . 2>&1
fi

echo "🎯 Build debug completed!"
