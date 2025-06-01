#!/bin/bash

# Railway deployment script
echo "🚀 Deploying to Railway..."

# Make sure you're logged in to Railway
railway login

# Link to your Railway project (if not already linked)
# railway link

# Deploy the application
railway up

echo "✅ Deployment completed!"
echo "🔗 Check your Railway dashboard for deployment status"
