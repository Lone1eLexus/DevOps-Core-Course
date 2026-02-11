package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestHealthEndpoint tests GET /health returns 200 and correct structure
func TestHealthEndpoint(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    rr := httptest.NewRecorder()

    healthHandler(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    var data map[string]interface{}
    if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
        t.Fatalf("Failed to parse JSON response: %v", err)
    }

    if _, ok := data["status"]; !ok {
        t.Error("Expected 'status' field in response")
    }

    if statusVal, ok := data["status"].(string); !ok || statusVal != "healthy" {
        t.Errorf("Expected status 'healthy', got '%v'", data["status"])
    }

    if _, ok := data["timestamp"]; !ok {
        t.Error("Expected 'timestamp' field in response")
    }

    if _, ok := data["uptime_seconds"]; !ok {
        t.Error("Expected 'uptime_seconds' field in response")
    }

    if uptime, ok := data["uptime_seconds"].(float64); !ok || uptime < 0 {
        t.Errorf("Expected uptime_seconds to be non-negative integer, got '%v'", data["uptime_seconds"])
    }
}

// TestMainEndpoint tests GET / returns 200 and all required fields
func TestMainEndpoint(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    mainHandler(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, status)
    }

    var data map[string]interface{}
    if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
        t.Fatalf("Failed to parse JSON response: %v", err)
    }

    // Service section
    service, ok := data["service"].(map[string]interface{})
    if !ok {
        t.Error("Expected 'service' field in response")
    } else {
        if name, ok := service["name"].(string); !ok || name != "devops-info-service" {
            t.Errorf("Expected service name 'devops-info-service', got '%v'", service["name"])
        }

        if version, ok := service["version"].(string); !ok || version != "1.0.0" {
            t.Errorf("Expected service version '1.0.0', got '%v'", service["version"])
        }

        if framework, ok := service["framework"].(string); !ok {
            t.Error("Expected 'framework' field in service")
        } else if framework != "Go net/http" {
            t.Errorf("Expected framework 'Go net/http', got '%v'", framework)
        }
    }

    // System section
    system, ok := data["system"].(map[string]interface{})
    if !ok {
        t.Error("Expected 'system' field in response")
    } else {
        if _, ok := system["hostname"]; !ok {
            t.Error("Expected 'hostname' field in system")
        }

        if _, ok := system["platform"]; !ok {
            t.Error("Expected 'platform' field in system")
        }

        if _, ok := system["cpu_count"]; !ok {
            t.Error("Expected 'cpu_count' field in system")
        }

        if cpuCount, ok := system["cpu_count"].(float64); !ok || cpuCount < 1 {
            t.Errorf("Expected cpu_count to be integer >= 1, got '%v'", system["cpu_count"])
        }
    }

    // Runtime section
    runtime, ok := data["runtime"].(map[string]interface{})
    if !ok {
        t.Error("Expected 'runtime' field in response")
    } else {
        if _, ok := runtime["uptime_seconds"]; !ok {
            t.Error("Expected 'uptime_seconds' field in runtime")
        }

        if _, ok := runtime["uptime_human"]; !ok {
            t.Error("Expected 'uptime_human' field in runtime")
        }

        if uptime, ok := runtime["uptime_seconds"].(float64); !ok || uptime < 0 {
            t.Errorf("Expected uptime_seconds to be non-negative, got '%v'", runtime["uptime_seconds"])
        }
    }

    // Request section
    request, ok := data["request"].(map[string]interface{})
    if !ok {
        t.Error("Expected 'request' field in response")
    } else {
        if _, ok := request["client_ip"]; !ok {
            t.Error("Expected 'client_ip' field in request")
        }

        if _, ok := request["user_agent"]; !ok {
            t.Error("Expected 'user_agent' field in request")
        }

        if _, ok := request["method"]; !ok {
            t.Error("Expected 'method' field in request")
        }

        if method, ok := request["method"].(string); !ok || method != "GET" {
            t.Errorf("Expected method 'GET', got '%v'", request["method"])
        }
    }

    // Endpoints list
    endpoints, ok := data["endpoints"].([]interface{})
    if !ok {
        t.Error("Expected 'endpoints' field in response")
    } else {
        if len(endpoints) < 2 {
            t.Errorf("Expected at least 2 endpoints, got %d", len(endpoints))
        }

        // Check that / and /health are in endpoints
        foundRoot := false
        foundHealth := false

        for _, ep := range endpoints {
            if epMap, ok := ep.(map[string]interface{}); ok {
                if path, ok := epMap["path"].(string); ok {
                    if path == "/" {
                        foundRoot = true
                    }
                    if path == "/health" {
                        foundHealth = true
                    }
                }
            }
        }

        if !foundRoot {
            t.Error("Expected endpoint '/' in endpoints list")
        }

        if !foundHealth {
            t.Error("Expected endpoint '/health' in endpoints list")
        }
    }
}

// Test404Endpoint tests non-existing endpoint returns 404 with proper JSON
func Test404Endpoint(t *testing.T) {
    req := httptest.NewRequest("GET", "/does-not-exist", nil)
    rr := httptest.NewRecorder()

    notFoundHandler(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("Expected status %d, got %d", http.StatusNotFound, status)
    }

    var data map[string]interface{}
    if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
        t.Fatalf("Failed to parse JSON response: %v", err)
    }

    if _, ok := data["error"]; !ok {
        t.Error("Expected 'error' field in response")
    }

    if errorVal, ok := data["error"].(string); !ok || errorVal != "Not Found" {
        t.Errorf("Expected error 'Not Found', got '%v'", data["error"])
    }

    if _, ok := data["message"]; !ok {
        t.Error("Expected 'message' field in response")
    }
}