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

## Installation as docker

1. Clone the repository:

```bash
    git clone https://github.com/Lone1eLexus/DevOps-Core-Course.git
    cd app_go
```

2. Build docker image:

```bash
    docker build . -t devops-info-service-go
```

## Pulling image from DockerHub

```bash
    docker pull lehus1/devops-info-service-go:1.0.0
```

## Running via docker

```bash
    docker run -p 8000:8000 devops-info-service-go
```


## API Endpoints
- GET / - Service and system information
- GET /health - Health check

## Configuration:

| Variable | Default | Description |
|----------|----------|--------------------|
| HOST | 0.0.0.0 | IP-Address |
| PORT | 8000 | Listening port |

## Status

![Go CI/CD](https://github.com/Lone1eLexus/DevOps-Core-Course/actions/workflows/go-ci.yml/badge.svg)

[![Coverage Status](https://coveralls.io/repos/github/Lone1eLexus/DevOps-Core-Course/badge.svg?branch=lab03&flag=go-unit-tests)](https://coveralls.io/github/Lone1eLexus/DevOps-Core-Course?branch=lab03)