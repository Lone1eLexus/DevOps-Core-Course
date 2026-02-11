# Lab 2 — Docker Containerization

## Docker Hub

**Process:**

```bash
  docker build . -t devops-info-service

  docker run -p 8000:8000 devops-info-service

  docker tag devops-info-service lehus1/devops-info-service:1.0.0

  docker push lehus1/devops-info-service:1.0.0
  The push refers to repository [docker.io/lehus1/devops-info-service]
  26796de0c7d5: Pushed 
  93465ce6a98b: Pushed 
  5b1a7b293d2f: Pushed 
  eb03c57d46ad: Pushed 
  44841c9dcdaa: Pushed 
  a915d0aa80cd: Pushed 
  ad1b18dd62d2: Pushed 
  d85cc8d16465: Pushed 
  e50a58335e13: Pushed 
  1.0.0: digest: sha256:60c6f77b2c48cf1042870134b5745cf771d0249d443aaf8c6262b4601c7a662a size: 2200
```

**URL:** https://hub.docker.com/repository/docker/lehus1/devops-info-service/general

**Strategy:**

  I chose Semantic because I wanted to choose it. (Also you can track changes thanks to this tyoe of tags)

## Docker Best Practices Applied

1. Non-Root User Implementation: Running as non-root reduces attack surface if container is compromised.

```bash
  RUN useradd --create-home --shell /bin/bash appuser 
  ...
  USER appuser
```

2. Specific Base Image Version: Same image version = same behavior everywhere. (And slim version, that needs less resourses)

```bash
  FROM python:3.13-slim
```

3. Layer Caching Optimization: Faster Builds, Efficient Development

```bash
  COPY requirements.txt .
  RUN pip install --no-cache-dir -r requirements.txt
```

4. Explicit Port Declaration: Clearly shows which port the app uses

```bash
  EXPOSE 8000
```


## Image Information & Decisions

### Base Image Choice: python:3.13-slim

### Reasons:

- Size Balance
- Compatibility: Based on Debian, works with most Python packages
- Security: Regular updates from Python/Debian teams
- Tooling: Includes pip and basic utilities



![alt text](../../app_go/docs/screenshots/size_com.png)

### Layers:

1. FROM python:3.13-slim

2. RUN useradd --create-home --shell /bin/bash appuser 

3. WORKDIR /app

4. COPY requirements.txt .

5. RUN pip install --no-cache-dir -r requirements.txt

6. COPY app.py .

7. USER appuser

8. EXPOSE 8000

9. CMD ["python", "app.py"]

### Explanation: 

All layers that change frequently come at the end, so we dont need to rebuild previous layers


## Build & Run Process

```bash
$ docker build . -t devops-info-service
[+] Building 1.1s (11/11) FINISHED                                                                                                                   docker:default
 => [internal] load build definition from dockerfile                                                                                                           0.0s
 => => transferring dockerfile: 273B                                                                                                                           0.0s
 => [internal] load metadata for docker.io/library/python:3.13-slim                                                                                            0.7s
 => [internal] load .dockerignore                                                                                                                              0.0s
 => => transferring context: 360B                                                                                                                              0.0s
 => [1/6] FROM docker.io/library/python:3.13-slim@sha256:51e1a0a317fdb6e170dc791bbeae63fac5272c82f43958ef74a34e170c6f8b18                                      0.0s
 => [internal] load build context                                                                                                                              0.0s
 => => transferring context: 63B                                                                                                                               0.0s
 => CACHED [2/6] RUN useradd --create-home --shell /bin/bash appuser                                                                                           0.0s
 => CACHED [3/6] WORKDIR /app                                                                                                                                  0.0s
 => CACHED [4/6] COPY requirements.txt .                                                                                                                       0.0s
 => CACHED [5/6] RUN pip install --no-cache-dir -r requirements.txt                                                                                            0.0s
 => CACHED [6/6] COPY app.py .                                                                                                                                 0.0s
 => exporting to image                                                                                                                                         0.1s
 => => exporting layers                                                                                                                                        0.0s
 => => writing image sha256:c96b2b7157e4ad8fecc0611acd5c3c515585380912331c9821c405b58dd7f8a6                                                                   0.0s
 => => naming to docker.io/library/devops-info-service   

$ docker run -p 8000:8000 devops-info-service
2026-02-03 21:22:50,305 - __main__ - INFO - Application starting...
2026-02-03 21:23:06,606 - __main__ - INFO - Incoming request: GET /health
2026-02-03 21:23:12,118 - __main__ - INFO - Incoming request: GET /
 ```

 ```bash
$ curl -s http://127.0.0.1:8000/health | jq '.'
{
  "status": "healthy",
  "timestamp": "2026-02-03T21:23:06.606131+00:00",
  "uptime_seconds": 16
}
$ curl -s http://127.0.0.1:8000/ | jq '.'
{
  "service": {
    "name": "devops-info-service",
    "version": "1.0.0",
    "description": "DevOps course info service",
    "framework": "FastAPI"
  },
  "system": {
    "hostname": "81b95dc31d3f",
    "platform": "Linux",
    "platform_version": "#36~24.04.1-Ubuntu SMP PREEMPT_DYNAMIC Wed Oct 15 15:45:17 UTC 2",
    "architecture": "x86_64",
    "cpu_count": 16,
    "python_version": "3.13.11"
  },
  "runtime": {
    "uptime_seconds": 21,
    "uptime_human": "0 hours, 0 minutes",
    "current_time": "2026-02-03T21:23:12.119433+00:00",
    "timezone": "UTC"
  },
  "request": {
    "client_ip": "172.17.0.1",
    "user_agent": "curl/8.5.0",
    "method": "GET",
    "path": "/"
  },
  "endpoints": [
    {
      "path": "/",
      "method": "GET",
      "description": "Service information"
    },
    {
      "path": "/health",
      "method": "GET",
      "description": "Health check"
    },
    {
      "path": "/docs",
      "method": "GET",
      "description": "Auto-generated API documentation"
    }
  ]
}
$ curl -s http://127.0.0.1:8000/docs

    <!DOCTYPE html>
    <html>
    <head>
    <link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css">
    <link rel="shortcut icon" href="https://fastapi.tiangolo.com/img/favicon.png">
    <title>Simple python app - Swagger UI</title>
    </head>
    <body>
    <div id="swagger-ui">
    </div>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <!-- `SwaggerUIBundle` is now available on the page -->
    <script>
    const ui = SwaggerUIBundle({
        url: '/openapi.json',
    "dom_id": "#swagger-ui",
"layout": "BaseLayout",
"deepLinking": true,
"showExtensions": true,
"showCommonExtensions": true,
oauth2RedirectUrl: window.location.origin + '/docs/oauth2-redirect',
    presets: [
        SwaggerUIBundle.presets.apis,
        SwaggerUIBundle.SwaggerUIStandalonePreset
        ],
    })
    </script>
    </body>
    </html>
$ curl -s http://127.0.0.1:8000/doc
{"error":"Not Found","message":"Endpoint /doc does not exist"}
```



## Technical Analysis

- Dockerfile uses optimal layer ordering. If we change it we will need to rebuild some layers.

- No sensitive data in layers: .dockerignore excludes .env, secrets, and reduce size

- Minimal attack surface: slim image has 85% fewer packages than full image

- Copy only necessary files (fans you can insert where you want)


## Challenges & Solutions

- Remembering the password DockerHub




  

