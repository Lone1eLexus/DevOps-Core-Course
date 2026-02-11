# DevOps Info Service

## Overview
A lightweight web service that provides real time system information and health status.

## Prerequisites
- Python 3.11+
- pip (Python package manager)

## Installation locally

1. Clone the repository:

```bash
    git clone https://github.com/Lone1eLexus/DevOps-Core-Course.git
    cd app_python
```

2. Create and activate virtual environment:

```bash
    python -m venv venv
    source venv/bin/activate  # Windows: venv\Scripts\activate
```

3. Install dependencies:

```bash
    pip install -r requirements.txt
```

## Running locally

- Basic Usage:
```bash
    python app.py
```

- With Custom Configuration
```bash
    PORT=8080 python app.py
    HOST=127.0.0.1 PORT=8080 python app.py
    DEBUG=True python app.py
```

## Installation as docker

1. Clone the repository:

```bash
    git clone https://github.com/Lone1eLexus/DevOps-Core-Course.git
    cd app_python
```

2. Build docker image:

```bash
    docker build . -t devops-info-service
```

## Pulling image from DockerHub

```bash
    docker pull lehus1/devops-info-service:1.0.0
```

## Running via docker

```bash
    docker run -p 8000:8000 devops-info-service
```



## API Endpoints
- GET / - Service and system information
- GET /health - Health check
- GET /docs - Auto-generated API documentation (FastAPI feature)

## Configuration:

| Variable | Default | Description |
|----------|----------|--------------------|
| HOST | 0.0.0.0 | IP-Address |
| PORT | 8000 | Listening port |
| DEBUG | False | Enable debug logging |