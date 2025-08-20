# Docker Setup Guide

This guide explains how to set up and run Jarvis with Docker Compose, including PostgreSQL and Temporal services.

## Prerequisites

- Docker
- Docker Compose
- Go 1.24.1 or later

## Quick Start

1. **Copy the environment file:**
   ```bash
   cp .env.example .env
   ```

2. **Update your API keys in `.env`:**
   - Set your `GEMINI_API_KEY` to your actual Google Gemini API key

3. **Start all services:**
   ```bash
   docker-compose up -d
   ```

4. **Verify services are running:**
   ```bash
   docker-compose ps
   ```

## Services Overview

### PostgreSQL (Port 4432)
- **Main Database:** `jarvis` 
- **Temporal Database:** `temporal`
- **User:** `jarvis`
- **Password:** `jarvis_password`
- **Connection Strings:** 
  - Main: `postgres://jarvis:jarvis_password@localhost:4432/jarvis`
  - Temporal: `postgres://jarvis:jarvis_password@localhost:4432/temporal`

### Temporal Server (Port 6233)
- **gRPC Endpoint:** `localhost:6233`
- **Used for:** Workflow orchestration and management

### Temporal Web UI (Port 7080)
- **URL:** http://localhost:7080
- **Used for:** Monitoring workflows, viewing history, debugging

## Development Workflow

### Starting Services
```bash
# Start all services in background
docker-compose up -d

# Start with logs visible
docker-compose up

# Start specific services
docker-compose up postgres temporal
```

### Stopping Services
```bash
# Stop all services
docker-compose down

# Stop and remove volumes (WARNING: This deletes all data!)
docker-compose down -v
```

### Viewing Logs
```bash
# View logs for all services
docker-compose logs

# View logs for specific service
docker-compose logs postgres
docker-compose logs temporal

# Follow logs in real-time
docker-compose logs -f
```

### Database Access

#### Connect to Main PostgreSQL Database
```bash
# Using psql
psql postgres://jarvis:jarvis_password@localhost:4432/jarvis

# Using Docker
docker-compose exec postgres psql -U jarvis -d jarvis
```

#### Connect to Temporal Database
```bash
# Using psql  
psql postgres://jarvis:jarvis_password@localhost:4432/temporal

# Using Docker
docker-compose exec postgres psql -U jarvis -d temporal
```

### Temporal Workflow Management

#### Access Temporal Web UI
Open http://localhost:7080 in your browser to:
- View running workflows
- Monitor task queues
- Debug workflow executions
- Browse workflow history

#### Using Temporal CLI (tctl)
```bash
# Install tctl (Temporal CLI)
go install go.temporal.io/server/tools/tctl@latest

# List workflows
tctl --address localhost:6233 workflow list

# Describe a workflow
tctl --address localhost:6233 workflow describe -w <workflow_id>
```

## Troubleshooting

### Common Issues

1. **Port Conflicts:**
   - Check if ports 4432, 6379, 6233, or 7080 are already in use
   - Modify ports in `docker-compose.yml` if needed

2. **Database Connection Issues:**
   - Ensure PostgreSQL containers are healthy: `docker-compose ps`
   - Check logs: `docker-compose logs postgres`

3. **Temporal Not Starting:**
   - Temporal depends on PostgreSQL being healthy
   - Check temporal logs: `docker-compose logs temporal`
   - Verify postgres is running: `docker-compose logs postgres`

### Health Checks
```bash
# Check all service health
docker-compose ps

# Test database connections
docker-compose exec postgres pg_isready -U jarvis -d jarvis
docker-compose exec postgres pg_isready -U jarvis -d temporal

# Test Temporal connection
docker-compose exec temporal tctl --address temporal:7233 cluster health
```

### Data Persistence

All data is persisted in Docker volumes:
- `postgres_data` - Main application database (includes both jarvis and temporal databases)
```

## Integration with Go Application

To use these services in your Go application, use the connection details from `.env`:

```go
// PostgreSQL connection
db, err := sql.Open("postgres", "postgres://jarvis:jarvis_password@localhost:4432/jarvis")

// Temporal client
client, err := temporal_client.Dial(temporal_client.Options{
    HostPort: "localhost:6233",
})
```
