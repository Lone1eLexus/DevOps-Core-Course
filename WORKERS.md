# Lab 17 — Cloudflare Workers Edge Deployment

## Task 1 — Cloudflare Setup

![alt text](image-1.png)

My v*n works bad on Ubuntu, so I use WSL now, but it has its own problems.

Later I will just copy this dir to my rep.

```bash
$ cd /mnt/e/DevOps-Course/
$ npm create cloudflare@latest -- edge-api

> npx
> create-cloudflare edge-api


──────────────────────────────────────────────────────────────────────────────────────────────────────────
👋 Welcome to create-cloudflare v2.68.2!
🧡 Let's get started.
📊 Cloudflare collects telemetry about your usage of Create-Cloudflare.

Learn more at: https://github.com/cloudflare/workers-sdk/blob/main/packages/create-cloudflare/telemetry.md
──────────────────────────────────────────────────────────────────────────────────────────────────────────

╭ Create an application with Cloudflare Step 1 of 3
│
├ In which directory do you want to create your application?
│ dir ./edge-api
│
├ What would you like to start with?
│ category Hello World example
│
├ Which template would you like to use?
│ type Worker only
│
├ Which language do you want to use?
│ lang TypeScript
│
├ Copying template files
│ files copied to project directory
│
├ Updating name in `package.json`
│ updated `package.json`
│
├ Installing dependencies
│ installed via `npm install`
│
├ Do you want to add an AGENTS.md file to help AI coding tools understand Cloudflare APIs?
│ yes agents
│
╰ Application created

╭ Configuring your application for Cloudflare Step 2 of 3
│
├ Installing wrangler A command line tool for building Cloudflare Workers
│ installed via `npm install wrangler --save-dev`
│
├ Retrieving current workerd compatibility date
│ compatibility date 2026-05-14
│
├ Generating types for your application
│ generated to `./worker-configuration.d.ts` via `npm run cf-typegen`
│
├ Installing @types/node
│ installed via npm
│
├ Do you want to use git for version control?
│ yes git
│
├ Initializing git repo
│ initialized git
│
├ Committing new files
│ git commit
│
╰ Application configured 

╭ Deploy with Cloudflare Step 3 of 3
│
├ Do you want to deploy your application?
│ no deploy via `npm run deploy`
│
╰ Done

────────────────────────────────────────────────────────────
🎉  SUCCESS  Application created successfully!

💻 Continue Developing
Change directories: cd edge-api
Deploy: npm run deploy

📖 Explore Documentation
https://developers.cloudflare.com/workers

🐛 Report an Issue
https://github.com/cloudflare/workers-sdk/issues/new/choose

💬 Join our Community
https://discord.cloudflare.com
────────────────────────────────────────────────────────────

$ cd edge-api
$ npx wrangler login

 ⛅️ wrangler 4.91.0
───────────────────
Attempting to login via OAuth...
Opening a link in your default browser: ...
Successfully logged in.
$ npx wrangler whoami

 ⛅️ wrangler 4.91.0
───────────────────
Getting User settings...
👋 You are logged in with an OAuth Token, associated with the email ....
┌─────────────────────────────────┬──────────────────────────────────┐
│ Account Name                    │ Account ID                       │
├─────────────────────────────────┼──────────────────────────────────┤
│ ...'s Account                   │ ...                              │
└─────────────────────────────────┴──────────────────────────────────┘
🔓 Token Permissions:
Scope (Access)
- account (read)
- user (read)
- workers (write)
- workers_kv (write)
- workers_routes (write)
- workers_scripts (write)
- workers_tail (read)
- d1 (write)
- pages (write)
- zone (read)
- ssl_certs (write)
- ai (write)
- ai-search (write)
- ai-search (run)
- queues (write)
- pipelines (write)
- secrets_store (write)
- artifacts (write)
- flagship (write)
- containers (write)
- cloudchamber (write)
- connectivity (admin)
- email_routing (write)
- email_sending (write)
- browser (write)
- offline_access
```

### Understanding Platform Concepts

| Concept | Description |
|---------|-------------|
| **Worker** | JavaScript/TypeScript function running on Cloudflare’s edge network. |
| **`workers.dev`** | Free subdomain for public Worker URLs. |
| **Bindings** | Resources injected into the Worker (vars, secrets, KV namespaces). |
| **Edge Runtime** | V8‑based runtime with `fetch` and standard Web APIs. |
| **Wrangler** | CLI for development, deployment, and configuration. |
| **Logs** | `console.log()` output visible via `wrangler tail` or dashboard. |

## Task 2 — Build and Deploy a Worker API

```bash
$ npx wrangler dev

 ⛅️ wrangler 4.91.0
───────────────────
╭───────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│  [b] open a browser [d] open devtools [e] open local explorer [t] start tunnel [c] clear console [x] to exit  │
╰───────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
⎔ Starting local server...
[wrangler:info] Ready on http://127.0.0.1:8787
[wrangler:info] GET / 200 OK (20ms)
[wrangler:info] GET /favicon.ico 404 Not Found (4ms)
[wrangler:info] GET / 200 OK (4ms)
[wrangler:info] GET / 200 OK (3ms)
[wrangler:info] GET /health 200 OK (8ms)
[wrangler:info] GET /health 200 OK (3ms)
[wrangler:info] GET / 200 OK (3ms)
[wrangler:info] GET /edge 200 OK (2ms)
[wrangler:info] GET /pohpoh 404 Not Found (3ms)
[wrangler:info] GET /non 404 Not Found (3ms)
[wrangler:info] GET / 200 OK (3ms)
[wrangler:info] GET /pohpoh 404 Not Found (3ms)
⎔ Shutting down local server...
```

```bash
$ curl http://localhost:8787/health -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"status":"ok","timestamp":"2026-05-14T11:44:17.047Z"}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 54
                    Content-Type: application/json

                    {"status":"ok","timestamp":"2026-05-14T11:44:17.047Z"}
Forms             :
Headers           : {[Content-Length, 54], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 54
$ curl http://localhost:8787/ -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"app":"edge-api","message":"Hello from Cloudflare Workers","version":"1.0.0","timestamp":"2026-05-
                    14T11:44:23.938Z"}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 117
                    Content-Type: application/json

                    {"app":"edge-api","message":"Hello from Cloudflare Workers","version":"1.0.0","timestamp":"2026-05-
                    14T11:44:23.938Z"}
Forms             :
Headers           : {[Content-Length, 117], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 117



$ curl http://localhost:8787/edge -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"message":"Edge metadata will appear here in Task 3","colo":"FRA","country":"DE"}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 82
                    Content-Type: application/json

                    {"message":"Edge metadata will appear here in Task 3","colo":"FRA","country":"DE"}
Forms             :
Headers           : {[Content-Length, 82], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 82



$ curl http://localhost:8787/pohpoh -UseBasicParsing
curl : Not Found
строка:1 знак:1
+ curl http://localhost:8787/pohpoh -UseBasicParsing
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebExc
   eption
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```

```bash
$ npx wrangler deploy

 ⛅️ wrangler 4.91.0
───────────────────
Total Upload: 0.86 KiB / gzip: 0.42 KiB
Worker Startup Time: 4 ms
Uploaded edge-api (6.85 sec)
Deployed edge-api triggers (5.90 sec)
  https://edge-api.lone...xus.workers.dev
Current Version ID: ...
```

```bash
$ curl https://edge-api.lone...xus.workers.dev/health
StatusCode        : 200
StatusDescription : OK
Content           : {"status":"ok","timestamp":"2026-05-14T14:51:40.889Z"}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=zSbgDVSTJR8wnjRbjW76bnH8fH8OVROPPoPf873PaKslrjzRyNm...
Forms             : {}
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=zSbgDVSTJR8wnjRbjW76bnH8fH8OVROPPoPf873PaKslrjzRyNmU6IEhaYWfkb
                    31ARDE%2BCJPHjc9Y6ERHmqU1dK7noUJ0iNsWFpU9h6nMPxceacJQPL5S3cWUuiAbDE5c0Ml95MMsU4DrboPF7yrcM7OmA%3D%3
                    D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY, 9fbab8ec8a6
                    edc94-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 54
```

## Task 3 — Global Edge Behavior

```bash
$ npx wrangler dev

 ⛅️ wrangler 4.91.0
───────────────────
╭───────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│  [b] open a browser [d] open devtools [e] open local explorer [t] start tunnel [c] clear console [x] to exit  │
╰───────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
⎔ Starting local server...
[wrangler:info] Ready on http://127.0.0.1:8787
[wrangler:info] GET /edge 200 OK (11ms)
⎔ Shutting down local server...
```

```bash
curl http://localhost:8787/edge -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"colo":"FRA","country":"DE","city":"Frankfurt am Main","asn":215439,"httpProtocol":"HTTP/1.1","tls
                    Version":"TLSv1.3","timestamp":"2026-05-14T12:13:51.626Z"}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 157
                    Content-Type: application/json

                    {"colo":"FRA","country":"DE","city":"Frankfurt am Main","asn":215439,"httpProtocol":"HTTP/1.1","tls
                    Version":"TLSv1.3","timestamp...
Forms             :
Headers           : {[Content-Length, 157], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 157
```

```bash
$ npx wrangler deploy

 ⛅️ wrangler 4.91.0
───────────────────
Total Upload: 1.18 KiB / gzip: 0.49 KiB
Worker Startup Time: 4 ms
Uploaded edge-api (11.09 sec)
Deployed edge-api triggers (6.08 sec)
  https://edge-api.lone...xus.workers.dev
Current Version ID: ...
```

```bash
$ curl https://edge-api.lone...xus.workers.dev/edge -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"colo":"FRA","country":"DE","city":"Frankfurt am Main","asn":215439,"httpProtocol":"HTTP/1.1","tls
                    Version":"TLSv1.3","timestamp":"2026-05-14T15:18:28.045Z"}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=kX1DjIJFxWP5VHJwIrPSXmEMAIhiBCBQVo2579%2F7AeLqi0e7N...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=kX1DjIJFxWP5VHJwIrPSXmEMAIhiBCBQVo2579%2F7AeLqi0e7NBiv72kcjcUc
                    JHflpP0XSgeWzX6BrQAYPD1mk5ibCCr1LjrdQQPeKbfQiAVnG1IJrt05Df3upCQi3mmuI3x0UhMAfpkEqHq838lwVRsX%2BA%3D
                    %3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY, 9fbae0293
                    ab5d3a0-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 157
```

### Edge Metadata Endpoint
Added `/edge` returning:
- `colo`: `{{ colo value }}`
- `country`: `{{ country value }}`
- `city`: `{{ city value }}`
- `asn`: `{{ asn value }}`
- `httpProtocol`: `{{ httpProtocol value }}`
- `tlsVersion`: `{{ tlsVersion value }}`

Workers run on Cloudflare’s edge network (300+ cities worldwide). When you deploy, the code is distributed to every data center automatically. There is **no manual “deploy to region X” step** – the Worker is available everywhere instantly.

**Key differences from VM/PaaS:**

| Aspect | Traditional VM/PaaS | Cloudflare Workers |
|--------|---------------------|--------------------|
| Region selection | Manual (e.g., us-east-1, eu-west-1) | Automatic global |
| Cold starts | Possible (scaling from zero) | Sub‑millisecond |
| Proximity to user | Only near chosen region | Always nearest edge |
| Deployment model | Push to one region, replicate manually | Single push → global instantly |


**Three ways to route traffic to a Worker:**

| Method | Description | Use case |
|--------|-------------|----------|
| **`workers.dev`** | Free subdomain (`<worker>.<your-sub>.workers.dev`) | Development, testing, quick demos |
| **Routes** | Attach Worker to an existing Cloudflare‑managed domain (e.g., `api.example.com/*`) | Production for custom domain |
| **Custom Domains** | Directly assign a domain to a Worker without a zone | Simple, but less flexible than Routes |

## Task 4 — Configuration, Secrets & Persistence

```bash
$ npx wrangler secret put API_TOKEN

 ⛅️ wrangler 4.91.0
───────────────────
√ Enter a secret value: ... ****
🌀 Creating the secret for the Worker "edge-api"
✨ Success! Uploaded secret API_TOKEN
$ npx wrangler secret put ADMIN_EMAIL

 ⛅️ wrangler 4.91.0
───────────────────
√ Enter a secret value: ... ******
🌀 Creating the secret for the Worker "edge-api"
✨ Success! Uploaded secret ADMIN_EMAIL
$ npx wrangler secret list
[
  {
    "name": "ADMIN_EMAIL",
    "type": "secret_text"
  },
  {
    "name": "API_TOKEN",
    "type": "secret_text"
  }
]
```

```bash
$ npx wrangler kv namespace create SETTINGS

 ⛅️ wrangler 4.91.0
───────────────────
Resource location: remote 

🌀 Creating namespace with title "SETTINGS"
✨ Success!
To access your new KV Namespace in your Worker, add the following snippet to your configuration file:
{
  "kv_namespaces": [
    {
      "binding": "SETTINGS",
      "id": "..."
    }
  ]
}
√ Would you like Wrangler to add it on your behalf? ... yes
√ What binding name would you like to use? ... SETTINGS
√ For local dev, do you want to connect to the remote resource instead of a local resource? ... no
```

```bash
$ curl http://localhost:8787/config -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"appName":"edge-api","courseName":"devops-core"} # Strange, but OK
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 49
                    Content-Type: application/json

                    {"appName":"edge-api","courseName":"devops-core"}
Forms             :
Headers           : {[Content-Length, 49], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 49



$ curl http://localhost:8787/counter -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"visits":1}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 12
                    Content-Type: application/json

                    {"visits":1}
Forms             :
Headers           : {[Content-Length, 12], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 12



$ curl http://localhost:8787/counter -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"visits":2}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 12
                    Content-Type: application/json

                    {"visits":2}
Forms             :
Headers           : {[Content-Length, 12], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 12

```

```bash
$ npx wrangler dev

 ⛅️ wrangler 4.91.0
───────────────────
Your Worker has access to the following bindings:
Binding                                                 Resource                  Mode
env.SETTINGS (383e6588e4984f388eb3238ad55dd709)         KV Namespace              local
env.APP_NAME ("edge-api")                               Environment Variable      local
env.COURSE_NAME ("devops-core")                         Environment Variable      local

❓ Your types might be out of date. Re-run `wrangler types` to ensure your types are correct.
╭───────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│  [b] open a browser [d] open devtools [e] open local explorer [t] start tunnel [c] clear console [x] to exit  │
╰───────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
⎔ Starting local server...
[wrangler:info] Ready on http://127.0.0.1:8787
[wrangler:info] GET /config 200 OK (8ms)
⎔ Reloading local server...
⎔ Local server updated and ready
⎔ Reloading local server...
⎔ Local server updated and ready
[wrangler:info] GET /config 200 OK (7ms)
[wrangler:info] GET /counter 200 OK (1086ms)
[wrangler:info] GET /counter 200 OK (27ms)
⎔ Shutting down local server...
$ npx wrangler dev

 ⛅️ wrangler 4.91.0
───────────────────
Your Worker has access to the following bindings:
Binding                                                 Resource                  Mode
env.SETTINGS (383e6588e4984f388eb3238ad55dd709)         KV Namespace              local
env.APP_NAME ("edge-api")                               Environment Variable      local
env.COURSE_NAME ("devops-core")                         Environment Variable      local

❓ Your types might be out of date. Re-run `wrangler types` to ensure your types are correct.
╭───────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│  [b] open a browser [d] open devtools [e] open local explorer [t] start tunnel [c] clear console [x] to exit  │
╰───────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
⎔ Starting local server...
[wrangler:info] Ready on http://127.0.0.1:8787
[wrangler:info] GET /counter 200 OK (1070ms)
⎔ Shutting down local server...
```

```bash
$  curl http://localhost:8787/counter -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"visits":3}
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 12
                    Content-Type: application/json

                    {"visits":3}
Forms             :
Headers           : {[Content-Length, 12], [Content-Type, application/json]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 12
```

## Task 5 — Observability & Operations

```bash
$ npx wrangler deploy

 ⛅️ wrangler 4.91.0
───────────────────
Total Upload: 1.82 KiB / gzip: 0.70 KiB
Worker Startup Time: 4 ms
Your Worker has access to the following bindings:
Binding                                                 Resource
env.SETTINGS (383e6588e4984f388eb3238ad55dd709)         KV Namespace
env.APP_NAME ("edge-api")                               Environment Variable
env.COURSE_NAME ("devops-core")                         Environment Variable

Uploaded edge-api (11.62 sec)
Deployed edge-api triggers (6.01 sec)
  https://edge-api.lone...xus.workers.dev
Current Version ID: ...
```

```bash
$ curl https://edge-api.lone...xus.workers.dev/counter -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"visits":3}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=gAKW4Gc%2BkeeV%2Fftk8%2Fg2GwTGIvq%2FuBFY%2FPhQQ1Jsn...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=gAKW4Gc%2BkeeV%2Fftk8%2Fg2GwTGIvq%2FuBFY%2FPhQQ1JsnNho0FWBxa73
                    rdT%2BUAiQRT%2F5o2ZX7GedyNgT5nmyxZzoUpKwA2dmCS8%2F5NbrhqxLGxdXE4TlYqK7IURbfBGFzJoL46Tfrsbwjcj%2BSFe
                    seMBmP%2BaSGA%3D%3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [C
                    F-RAY, 9fbb2b5a8be5d351-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 12



$ curl https://edge-api.lone...xus.workers.dev/counter -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"visits":4}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=Y87BbuGzETnavKCAUwD9JlbamlZ%2B6jmZNlBtmbyYpy2IcxE5X...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=Y87BbuGzETnavKCAUwD9JlbamlZ%2B6jmZNlBtmbyYpy2IcxE5XNGOmE9vtnmN
                    nOVii8AiM1YCOZDpXBmGMHGErtv9eihOrCaeCM82J7Ptxvilv%2BfrObvlzo9w%2FOakAPm8MZ5x4JTe46fTkakHU5sqMG64tw%
                    3D%3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY, 9fbb2b7
                    50ed6d351-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 12



$ curl https://edge-api.lone...xus.workers.dev/ -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"app":"edge-api","message":"Hello from Cloudflare Workers","version":"1.0.0","timestamp":"2026-05-
                    14T16:10:00.843Z"}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=myHrQ7Dgx2dFhRAfaqf6xZ0aYZrFzPXYqpSJEljBsQ73zP2FxcC...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=myHrQ7Dgx2dFhRAfaqf6xZ0aYZrFzPXYqpSJEljBsQ73zP2FxcCOmaH6JFVArC
                    Cw4U1m%2BX6YLHZ0N%2FOKTQKe8OkWE6Y8u4XhBHQe3%2FpJ9rf9RUu1EG1Gh2LS19cT75z6Am1pWBEVA2BunMEBmUVt4F%2Fx%
                    2Bw%3D%3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY, 9fb
                    b2bab4902d351-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 117



$ curl https://edge-api.lone...xus.workers.dev/config -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"appName":"edge-api","courseName":"devops-core","ApiToken":"haha","adminEmail":"really"}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=8N6jHeNFY3%2BktlP1RkAaTxM72QXJ6N4AAH5iak6jrmiCKImpB...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=8N6jHeNFY3%2BktlP1RkAaTxM72QXJ6N4AAH5iak6jrmiCKImpBXCaJqBLX7la
                    JXNQWUwVrvkprzGM%2BeYGlyKlCioUbCzFzwJKbAt87vDE6vaqZ3sgw9FtKKiHJ9%2FbFcTYPXLHZxfl7d2oDeUT2nt5Hy4GRA%
                    3D%3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY, 9fbb2bc
                    cfc33d351-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 89



$ curl https://edge-api.lone...xus.workers.dev/edge -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"colo":"FRA","country":"DE","city":"Frankfurt am Main","asn":215439,"httpProtocol":"HTTP/1.1","tls
                    Version":"TLSv1.3","timestamp":"2026-05-14T16:10:10.274Z"}
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/rep
                    ort/v4?s=RDZ3r9oBTB53Qs8kgjhyL6vtDXeJLLO04XiRoSniE9b6frmeKzO...
Forms             :
Headers           : {[Connection, keep-alive], [Report-To, {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"http
                    s://a.nel.cloudflare.com/report/v4?s=RDZ3r9oBTB53Qs8kgjhyL6vtDXeJLLO04XiRoSniE9b6frmeKzOqTE0hLc%2B7
                    MPEZO%2BuW7vP7GAKXUPZyiCbjqls519YYpNO%2BPLtwwRp%2FUr76IY6kIy69mUGtMG%2B%2BSJw%2FSuODmsp8yMMUeIC0Rs7
                    EiCxolw%3D%3D"}]}], [Nel, {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}], [CF-RAY,
                     9fbb2be63ed1d351-FRA]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 157



$ curl https://edge-api.lone...xus.workers.dev/lll -UseBasicParsing
curl : Not Found
строка:1 знак:1
+ curl https://edge-api.lone...xus.workers.dev/lll -UseBasicParsing
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebExc
   eption
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```

```bash
$ npx wrangler tail

 ⛅️ wrangler 4.91.0
───────────────────
Successfully created tail, expires at 2026-05-14T22:07:57Z
Connected to edge-api, waiting for logs...
GET https://edge-api.lone...xus.workers.dev/counter - Ok @ 14.05.2026, 19:09:47
  (log) [2026-05-14T16:09:47.929Z] GET /counter - colo: FRA
GET https://edge-api.lone...xus.workers.dev/counter - Ok @ 14.05.2026, 19:09:52
  (log) [2026-05-14T16:09:52.171Z] GET /counter - colo: FRA
GET https://edge-api.lone...xus.workers.dev/ - Ok @ 14.05.2026, 19:10:00
  (log) [2026-05-14T16:10:00.843Z] GET / - colo: FRA
GET https://edge-api.lone...xus.workers.dev/config - Ok @ 14.05.2026, 19:10:06
  (log) [2026-05-14T16:10:06.240Z] GET /config - colo: FRA
GET https://edge-api.lone...xus.workers.dev/edge - Ok @ 14.05.2026, 19:10:10
  (log) [2026-05-14T16:10:10.274Z] GET /edge - colo: FRA
GET https://edge-api.lone...xus.workers.dev/lll - Ok @ 14.05.2026, 19:10:23
  (log) [2026-05-14T16:10:23.454Z] GET /lll - colo: FRA
```

```bash
$ npx wrangler deployments list

 ⛅️ wrangler 4.91.0
───────────────────
Created:     2026-05-14T14:47:24.522Z
Author:      ...
Source:      Upload
Message:     Automatic deployment on upload.
Version(s):  (100%) ee79e584-4df3-4b3b-9ecc-edbc4770274b
                 Created:  2026-05-14T14:47:24.522Z
                     Tag:  -
                 Message:  -

Created:     2026-05-14T15:15:20.348Z
Author:      ...
Source:      Unknown (deployment)
Message:     -
Version(s):  (100%) 4f40cc32-67b5-44c8-adbc-16d3282c5dad
                 Created:  2026-05-14T15:15:17.816Z
                     Tag:  -
                 Message:  -

Created:     2026-05-14T15:30:27.242Z
Author:      ...
Source:      Secret Change
Message:     -
Version(s):  (100%) 803848bb-ceae-4d9b-84ba-01470b464186
                 Created:  2026-05-14T15:30:27.242Z
                     Tag:  -
                 Message:  -

Created:     2026-05-14T15:30:43.545Z
Author:      ...
Source:      Secret Change
Message:     -
Version(s):  (100%) 44e211e5-878e-4e9a-9c99-810527b5959a
                 Created:  2026-05-14T15:30:43.545Z
                     Tag:  -
                 Message:  -

Created:     2026-05-14T15:42:01.515Z
Author:      ...
Source:      Unknown (deployment)
Message:     -
Version(s):  (100%) 9e2dcc4a-add9-46bd-b7bd-ff50c3996863
                 Created:  2026-05-14T15:41:58.195Z
                     Tag:  -
                 Message:  -

Created:     2026-05-14T16:05:18.019Z
Author:      ...
Source:      Unknown (deployment)
Message:     -
Version(s):  (100%) 985d5236-faee-423f-aee9-dbdff52ba17e
                 Created:  2026-05-14T16:05:15.165Z
                     Tag:  -
                 Message:  -
$ npx wrangler rollback

 ⛅️ wrangler 4.91.0
───────────────────
├ Your current deployment has 1 version(s):
│
│ (100%) 985d5236-faee-423f-aee9-dbdff52ba17e
│       Created:  2026-05-14T16:05:15.165761Z
│           Tag:  -
│       Message:  -
│
√ Please provide an optional message for this rollback (120 characters max) ... Rollback
│
├  WARNING  You are about to rollback to Worker Version 9e2dcc4a-add9-46bd-b7bd-ff50c3996863.
│ This will immediately replace the current deployment and become the active deployment across all your deployed triggers.
│ However, your local development environment will not be affected by this rollback.
│ Rolling back to a previous deployment will not rollback any of the bound resources (Durable Object, D1, R2, KV, etc).
│
│ (100%) 9e2dcc4a-add9-46bd-b7bd-ff50c3996863
│       Created:  2026-05-14T15:41:58.195779Z
│           Tag:  -
│       Message:  -
│
√ Are you sure you want to deploy this Worker Version to 100% of traffic? ... yes
Performing rollback...
│
╰  SUCCESS  Worker Version 9e2dcc4a-add9-46bd-b7bd-ff50c3996863 has been deployed to 100% of traffic.

Current Version ID: 9e2dcc4a-add9-46bd-b7bd-ff50c3996863
```

![alt text](image-3.png)

Metrics on dashboard are unavailable :(

## Task 6 — Documentation & Comparison

**Worker url:** https://edge-api.lone...xus.workers.dev, where ... = ele1 rotated 

### Kubernetes vs Cloudflare Workers Comparison

| Aspect | Kubernetes | Cloudflare Workers |
|--------|------------|--------------------|
| **Setup complexity** | High – cluster provisioning, networking, storage, Ingress | Low – account, CLI, deploy |
| **Deployment speed** | Minutes to hours (image build, push, rollout) | Seconds (single `wrangler deploy`) |
| **Global distribution** | Manual (multiple clusters, Geo‑DNS) | Automatic (300+ edge locations) |
| **Cost (for small apps)** | High – control plane, nodes, load balancers | Free tier (100k requests/day) or very low |
| **State/persistence model** | PVCs, StatefulSets, external databases | KV, D1, R2 (globally distributed) |
| **Control/flexibility** | Full – custom images, networking, security policies | Limited – runtime, dependencies, no raw TCP |
| **Best use case** | Long‑running containers, complex microservices, batch jobs | Low‑latency HTTP APIs, global edge logic, CDN‑integrated features |

---

### When to Use Each

| Scenarios favoring **Kubernetes** | Scenarios favoring **Cloudflare Workers** |
|-----------------------------------|-------------------------------------------|
| Stateful workloads (databases, message queues) | Stateless HTTP APIs |
| Machine learning model serving | Real‑time personalisation |
| Legacy apps that need full OS compatibility | Webhook handlers |
| Long‑running background tasks | Global request routing / edge logic |
| Custom networking (e.g., eBPF, service mesh) | Simple CRUD with KV / D1 |

**My recommendation:**  
- Use **Workers** for low‑latency, globally distributed APIs and lightweight logic.  
- Use **Kubernetes** for complex, stateful, or resource‑intensive workloads where you need full control over the runtime.
- Workers simple, but Kubernetes is better for complex apps.
---

### Reflection

#### What felt easier than Kubernetes?
- No cluster setup – just `npm create cloudflare` and deploy.
- Automatic HTTPS and global CDN out of the box.
- Simple CLI (`wrangler`) with built‑in secrets, KV, and logs.
- No need to manage Ingress or load balancers.

#### What felt more constrained?
- Limited runtime (no Node.js built‑ins like `fs` or `child_process`).
- Cold start is very fast but still slightly perceptible for infrequently called Workers.
- No raw WebSocket server (but you can use Durable Objects for stateful WebSockets).
- Persistent storage is eventually consistent (KV may be stale for a few seconds).

#### What changed because Workers is not a Docker host?
- Could not run my existing Docker‑based Python app; had to rewrite in TypeScript.
- No image building – just `wrangler deploy` pushes source code.
- Dependencies are bundled into the Worker (max 1MB after compression).
- No shell access – debugging means logs and `wrangler tail`.



