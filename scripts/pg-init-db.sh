#!/bin/bash

# This script sets up a PostgreSQL database with a new user and database.
# It requires the following environment variables to be set:
# DB_USER, DB_PASS, DB_NAME, POSTGRES_USER

set -e

log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $1"
}

# Validate input
if [ -z "$DB_USER" ] || [ -z "$DB_PASS" ] || [ -z "$DB_NAME" ] || [ -z "$POSTGRES_USER" ]; then
    log "Error: One or more required environment variables are not set."
    echo "Required variables: DB_USER, DB_PASS, DB_NAME, POSTGRES_USER"
    exit 1
fi

log "Starting database setup"

# Use PGPASSWORD for more secure password handling
export PGPASSWORD="$DB_PASS"

if ! psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" \
    -v db_user="$DB_USER" \
    -v db_pass="$DB_PASS" \
    -v db_name="$DB_NAME" <<-EOSQL
    CREATE USER :"db_user" WITH PASSWORD :'db_pass';
    CREATE DATABASE :"db_name";
    GRANT ALL PRIVILEGES ON DATABASE :"db_name" TO :"db_user";
    ALTER DATABASE :"db_name" OWNER TO :"db_user";
EOSQL
then
    log "Error: Failed to set up database"
    exit 1
fi

log "Database setup completed successfully"

# Unset PGPASSWORD for security
unset PGPASSWORD