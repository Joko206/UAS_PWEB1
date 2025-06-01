#!/bin/bash

echo "🐳 Testing Docker build locally..."

# Clean up any existing containers/images
echo "🧹 Cleaning up..."
docker stop brainquiz-test 2>/dev/null || true
docker rm brainquiz-test 2>/dev/null || true
docker rmi brainquiz-test 2>/dev/null || true

# Build the Docker image
echo "🔨 Building Docker image..."
docker build -t brainquiz-test .

if [ $? -ne 0 ]; then
    echo "❌ Docker build failed!"
    exit 1
fi

echo "✅ Docker build successful!"

# Test run the container
echo "🚀 Testing container run..."
docker run -d --name brainquiz-test -p 8000:8000 \
    -e DATABASE_URL="postgresql://test:test@localhost:5432/test?sslmode=disable" \
    brainquiz-test

# Wait a bit for startup
sleep 5

# Check if container is running
if docker ps | grep -q brainquiz-test; then
    echo "✅ Container is running!"
    
    # Test health endpoint
    echo "🏥 Testing health endpoint..."
    curl -f http://localhost:8000/health || echo "⚠️  Health check failed (expected if no DB)"
else
    echo "❌ Container failed to start!"
    echo "📋 Container logs:"
    docker logs brainquiz-test
fi

# Cleanup
echo "🧹 Cleaning up test container..."
docker stop brainquiz-test 2>/dev/null || true
docker rm brainquiz-test 2>/dev/null || true

echo "🎉 Docker test completed!"
