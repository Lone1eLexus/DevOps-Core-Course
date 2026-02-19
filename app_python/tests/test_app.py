"""
Unit tests for DevOps Info Service (FastAPI)
"""
import pytest
from fastapi.testclient import TestClient
from app import app

client = TestClient(app)

def test_health_endpoint():
    """Test GET /health returns 200 and correct structure"""
    response = client.get("/health")
    assert response.status_code == 200
    data = response.json()
    assert "status" in data
    assert data["status"] == "healthy"
    assert "timestamp" in data
    assert "uptime_seconds" in data
    assert isinstance(data["uptime_seconds"], int)

def test_main_endpoint():
    """Test GET / returns 200 and all required fields"""
    response = client.get("/")
    assert response.status_code == 200
    data = response.json()

    # Service section
    assert "service" in data
    svc = data["service"]
    assert svc["name"] == "devops-info-service"
    assert svc["version"] == "1.0.0"
    assert svc["framework"] in ["FastAPI"]

    # System section
    assert "system" in data
    sys = data["system"]
    assert "hostname" in sys
    assert "platform" in sys
    assert "cpu_count" in sys
    assert isinstance(sys["cpu_count"], int)

    # Runtime section
    assert "runtime" in data
    rt = data["runtime"]
    assert "uptime_seconds" in rt
    assert "uptime_human" in rt
    assert isinstance(rt["uptime_seconds"], int)

    # Request section
    assert "request" in data
    req = data["request"]
    assert "client_ip" in req
    assert "user_agent" in req
    assert "method" in req
    assert req["method"] == "GET"

    # Endpoints list
    assert "endpoints" in data
    endpoints = data["endpoints"]
    assert len(endpoints) >= 2
    paths = [ep["path"] for ep in endpoints]
    assert "/" in paths
    assert "/health" in paths

def test_404_endpoint():
    """Test non-existing endpoint returns 404 with proper JSON"""
    response = client.get("/does-not-exist")
    assert response.status_code == 404
    data = response.json()
    assert "error" in data
    assert data["error"] == "Not Found"
    assert "message" in data