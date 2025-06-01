#!/bin/bash

echo "🚀 Starting BrainQuiz API..."

# Debug information
echo "📋 Environment Debug:"
echo "- Current directory: $(pwd)"
echo "- Files in directory:"
ls -la

# Check if main binary exists
if [ ! -f "./main" ]; then
    echo "❌ Binary './main' not found!"
    echo "🔨 Attempting to build..."
    
    # Try to build the binary
    if command -v go >/dev/null 2>&1; then
        echo "✅ Go found, building..."
        go build -o main .
        
        if [ -f "./main" ]; then
            echo "✅ Build successful!"
        else
            echo "❌ Build failed!"
            exit 1
        fi
    else
        echo "❌ Go not found and binary doesn't exist!"
        exit 1
    fi
fi

# Make sure binary is executable
chmod +x ./main

# Verify binary
echo "📊 Binary info:"
ls -la main
file main

# Start the application
echo "🚀 Starting application..."
exec ./main
