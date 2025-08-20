#!/bin/bash
set -e

# This script creates multiple databases for the application
# It's executed during PostgreSQL container initialization

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- Create temporal databases if they don't exist
    SELECT 'CREATE DATABASE temporal_jarvis' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'temporal_jarvis')\gexec
    SELECT 'CREATE DATABASE temporal_visibility_jarvis' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'temporal_visibility_jarvis')\gexec
    
    -- Grant all privileges on temporal databases to the main user
    GRANT ALL PRIVILEGES ON DATABASE temporal_jarvis TO $POSTGRES_USER;
    GRANT ALL PRIVILEGES ON DATABASE temporal_visibility_jarvis TO $POSTGRES_USER;
EOSQL

echo "Databases initialized successfully!"
echo "Created databases: jarvis, temporal_jarvis, temporal_visibility_jarvis"
echo "User '$POSTGRES_USER' has access to all databases"
