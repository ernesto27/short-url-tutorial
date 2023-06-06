env_file=".env"

# Check if the .env file exists
if [ -f "$env_file" ]; then
    # Load the variables into the current shell environment
    source "$env_file"
    echo "Environment variables loaded from $env_file:"
    
    zip -r app.zip . -x ".git/*"
    az webapp deploy --resource-group $AZURE_RESOURCE_GROUP --name $AZURE_NAME --src-path app.zip
else
    echo "Error: $env_file not found."
fi

