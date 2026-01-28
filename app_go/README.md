# DevOps Info Service

## Overview
A lightweight web service that provides real time system information and health status.

## Prerequisites
- Go 1.21 or higher

## Installation

1. Clone the repository:

```bash
   git clone https://github.com/Lone1eLexus/DevOps-Core-Course.git
   cd app_go
```

2. Initialize and build:

```bash
    go mod init
    go build -o app
```

## Running

- Basic Usage:
```bash
    ./app
```

- With Custom Configuration
```bash
    # Custom port
    PORT=9090 ./app

    # Specific host and port
    HOST=127.0.0.1 PORT=3000 ./app
```

## API Endpoints
- GET / - Service and system information
- GET /health - Health check

## Configuration:

| Variable | Default | Description |
|----------|----------|--------------------|
| HOST | 0.0.0.0 | IP-Address |
| PORT | 8000 | Listening port |