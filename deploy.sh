#!/bin/bash

# Railway deployment script
echo "🚀 Preparing deployment to Railway..."

# Clean up previous builds
echo "🧹 Cleaning up..."
rm -f main

# Test build locally first
echo "🔨 Testing local build..."
go mod tidy
go build -o main .

if [ $? -eq 0 ]; then
    echo "✅ Local build successful!"
    rm -f main
else
    echo "❌ Local build failed! Fix errors before deploying."
    exit 1
fi

echo "📋 Deployment checklist:"
echo "1. ✅ Go version: $(go version)"
echo "2. ✅ Build test: Passed"
echo "3. 🔧 Make sure these environment variables are set in Railway:"
echo "   - DATABASE_URL (PostgreSQL connection string)"
echo "   - PORT (should be auto-set by Railway)"
echo "   - CORS_ORIGINS (if needed)"
echo ""
echo "4. 🚀 Deploy using one of these methods:"
echo "   a) Railway CLI: railway up"
echo "   b) Git push to connected repository"
echo "   c) Railway dashboard manual deploy"
echo ""
echo "5. 🔍 After deployment, check:"
echo "   - Railway logs for startup messages"
echo "   - Health endpoint: https://your-app.railway.app/health"
echo ""
echo "📝 Common issues:"
echo "   - Database connection: Check DATABASE_URL format"
echo "   - Port binding: Railway auto-sets PORT variable"
echo "   - SSL mode: Use 'require' for Railway PostgreSQL"
