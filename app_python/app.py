"""
DevOps Info Service
Main application module providing system introspection APIs
"""
import os
import platform
import socket
from datetime import datetime, timezone

import logging

from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
import uvicorn


# config
HOST = os.getenv("HOST", "0.0.0.0")
PORT = int(os.getenv("PORT", "8000"))
DEBUG = os.getenv("DEBUG", "False").lower() == "true"

app = FastAPI(
    title="Simple python app"
)

# Logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

logger.info('Application starting...')

# System info


def get_system_info():
    """Collect system information."""
    try:
        return {
            "hostname": socket.gethostname(),
            "platform": platform.system(),
            "platform_version": platform.version(),
            "architecture": platform.machine(),
            "cpu_count": os.cpu_count() or 1,
            "python_version": platform.python_version()
        }
    except Exception as e:
        logger.error(f"Error getting system info: {e}")
        return {
            "hostname": "unknown",
            "platform": "unknown",
            "platform_version": "unknown",
            "architecture": "unknown",
            "cpu_count": 1,
            "python_version": "unknown"
        }


# Time
start_time = datetime.now(timezone.utc)


def get_uptime():
    delta = datetime.now(timezone.utc) - start_time
    seconds = int(delta.total_seconds())
    hours = seconds // 3600
    minutes = (seconds % 3600) // 60
    return {
        'seconds': seconds,
        'human': f"{hours} hours, {minutes} minutes"
    }

# Endpoints


@app.get("/health")
def health(request: Request):
    """Health endpoint - current health of server."""
    logger.info(f"Incoming request: {request.method} {request.url.path}")
    return {
        'status': 'healthy',
        'timestamp': datetime.now(timezone.utc).isoformat(),
        'uptime_seconds': get_uptime()['seconds']
    }


@app.get("/")
def index(request: Request):
    """Main endpoint - service and system information."""
    logger.info(f"Incoming request: {request.method} {request.url.path}")
    return {
        "service": {
            "name": "devops-info-service",
            "version": "1.0.0",
            "description": "DevOps course info service",
            "framework": "FastAPI"
        },
        "system": get_system_info(),
        "runtime": {
            "uptime_seconds": get_uptime()["seconds"],
            "uptime_human": get_uptime()["human"],
            "current_time": datetime.now(timezone.utc).isoformat(),
            "timezone": "UTC"
        },
        "request": {
            "client_ip": request.client.host,
            "user_agent": request.headers.get("user-agent"),
            "method": request.method,
            "path": request.url.path
        },
        "endpoints": [
            {"path": "/", "method": "GET", "description":
             "Service information"},
            {"path": "/health", "method": "GET", "description":
             "Health check"},
            {"path": "/docs", "method": "GET",
                "description": "Auto-generated API documentation"}
        ]
    }

# Error Handling


@app.exception_handler(404)
def not_found(request: Request, exc):
    """Handle 404 Not Found errors."""
    return JSONResponse(
        status_code=404,
        content={
            "error": "Not Found",
            "message": f"Endpoint {request.url.path} does not exist"
        }
    )


@app.exception_handler(500)
def internal_error(request: Request, exc):
    """Handle 500 Internal Server errors."""
    logger.error(f"Internal server error: {exc}")
    return JSONResponse(
        status_code=500,
        content={
            "error": "Internal Server Error",
            "message": "An unexpected error occurred"
        }
    )


# main


if __name__ == "__main__":
    uvicorn.run(
        app,
        host=HOST,
        port=PORT,
        log_level="info" if DEBUG else "warning")
