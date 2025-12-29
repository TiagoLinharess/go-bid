#!/bin/bash

echo "Waiting database to be ready..."
until pg_isready -h db -U postgres -d gobid; do
    sleep 1
done

echo "Creating database if dosen't exists..."
psql -h db -U postgres -c "CREATE DATABASE gobid;" || true

echo "Appling migrations..."
cd /app/internal/store/pgstore/migrations
tern migrate

echo "Starting application..."
cd /app

exec ./main