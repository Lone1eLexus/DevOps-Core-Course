package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "runtime"
    "strings"
    "time"
)

type Service struct {
    Name        string `json:"name"`
    Version     string `json:"version"`
    Description string `json:"description"`
    Framework   string `json:"framework"`
}

type System struct {
    Hostname        string `json:"hostname"`
    Platform        string `json:"platform"`
    PlatformVersion string `json:"platform_version"`
    Architecture    string `json:"architecture"`
    CPUCount        int    `json:"cpu_count"`
    GoVersion       string `json:"go_version"`
}

type Runtime struct {
    UptimeSeconds int    `json:"uptime_seconds"`
    UptimeHuman   string `json:"uptime_human"`
    CurrentTime   string `json:"current_time"`
    Timezone      string `json:"timezone"`
}

type Request struct {
    ClientIP  string `json:"client_ip"`
    UserAgent string `json:"user_agent"`
    Method    string `json:"method"`
    Path      string `json:"path"`
}

type Endpoint struct {
    Path        string `json:"path"`
    Method      string `json:"method"`
    Description string `json:"description"`
}

type ServiceInfo struct {
    Service   Service   `json:"service"`
    System    System    `json:"system"`
    Runtime   Runtime   `json:"runtime"`
    Request   Request   `json:"request"`
    Endpoints []Endpoint `json:"endpoints"`
}

type HealthResponse struct {
    Status        string `json:"status"`
    Timestamp     string `json:"timestamp"`
    UptimeSeconds int    `json:"uptime_seconds"`
}

var startTime = time.Now()

func getUptime() (int, string) {
    elapsed := time.Since(startTime)
    seconds := int(elapsed.Seconds())
    hours := seconds / 3600
    minutes := (seconds % 3600) / 60

    hourStr := "hour"
    if hours != 1 {
        hourStr += "s"
    }
    minuteStr := "minute"
    if minutes != 1 {
        minuteStr += "s"
    }

    return seconds, fmt.Sprintf("%d %s, %d %s", hours, hourStr, minutes, minuteStr)
}

func getHostname() string {
    hostname, err := os.Hostname()
    if err != nil {
        return "unknown"
    }
    return hostname
}

func getClientIP(r *http.Request) string {
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ips := strings.Split(forwarded, ",")
        return strings.TrimSpace(ips[0])
    }
    
    host, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return r.RemoteAddr
    }
    return host
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    uptimeSeconds, uptimeHuman := getUptime()

    info := ServiceInfo{
        Service: Service{
            Name:        "devops-info-service",
            Version:     "1.0.0",
            Description: "DevOps course info service",
            Framework:   "Go net/http",
        },
        System: System{
            Hostname:        getHostname(),
            Platform:        runtime.GOOS,
            PlatformVersion: runtime.GOOS + "/" + runtime.GOARCH,
            Architecture:    runtime.GOARCH,
            CPUCount:        runtime.NumCPU(),
            GoVersion:       runtime.Version(),
        },
        Runtime: Runtime{
            UptimeSeconds: uptimeSeconds,
            UptimeHuman:   uptimeHuman,
            CurrentTime:   time.Now().UTC().Format(time.RFC3339),
            Timezone:      "UTC",
        },
        Request: Request{
            ClientIP:  getClientIP(r),
            UserAgent: r.Header.Get("User-Agent"),
            Method:    r.Method,
            Path:      r.URL.Path,
        },
        Endpoints: []Endpoint{
            {Path: "/", Method: "GET", Description: "Service information"},
            {Path: "/health", Method: "GET", Description: "Health check"},
        },
    }

    jsonResponse(w, http.StatusOK, info)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    uptimeSeconds, _ := getUptime()

    response := HealthResponse{
        Status:        "healthy",
        Timestamp:     time.Now().UTC().Format(time.RFC3339),
        UptimeSeconds: uptimeSeconds,
    }

    jsonResponse(w, http.StatusOK, response)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{
        "error":   "Not Found",
        "message": fmt.Sprintf("Endpoint %s does not exist", r.URL.Path),
    }
    jsonResponse(w, http.StatusNotFound, response)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
    log.SetFlags(log.LstdFlags)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    host := os.Getenv("HOST")
    if host == "" {
        host = "0.0.0.0"
    }

    log.Printf("Starting DevOps Info Service on %s:%s", host, port)

    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            mainHandler(w, r)
        } else {
            notFoundHandler(w, r)
        }
    })
    mux.HandleFunc("/health", loggingMiddleware(healthHandler))

    addr := host + ":" + port
    log.Printf("Listening on %s", addr)
    
    err := http.ListenAndServe(addr, mux)
    if err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}