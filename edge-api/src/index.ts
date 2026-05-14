export interface Env {
  APP_NAME?: string;
  COURSE_NAME?: string;
  API_TOKEN?: string;
  ADMIN_EMAIL?: string;
  SETTINGS?: KVNamespace;   // will be added later
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);

    console.log(`[${new Date().toISOString()}] ${request.method} ${url.pathname} - colo: ${request.cf?.colo ?? 'unknown'}`);

    // Health check endpoint
    if (url.pathname === "/health") {
      return Response.json({ status: "ok", timestamp: new Date().toISOString() });
    }

    // Root endpoint – general app info
    if (url.pathname === "/") {
      return Response.json({
        app: env.APP_NAME || "edge-api",
        message: "Hello from Cloudflare Workers",
        version: "1.0.0",
        timestamp: new Date().toISOString(),
      });
    }

    // Edge metadata endpoint (placeholder for Task 3)
    if (url.pathname === "/edge") {
      return Response.json({
        // Required fields
        colo: request.cf?.colo || "unknown",
        country: request.cf?.country || "unknown",
        // Additional fields (choose at least one)
        city: request.cf?.city || "unknown",
        asn: request.cf?.asn || "unknown",
        httpProtocol: request.cf?.httpProtocol || "unknown",
        tlsVersion: request.cf?.tlsVersion || "unknown",
        // Optional: include request time
        timestamp: new Date().toISOString(),
      });
    }

    if (url.pathname === "/config") {
      return Response.json({
        appName: env.APP_NAME,
        courseName: env.COURSE_NAME,
        ApiToken: env.API_TOKEN,
        adminEmail: env.ADMIN_EMAIL,
      });
    }

    if (url.pathname === "/counter") {
      const key = "visits";
      let visits = await env.SETTINGS?.get(key);
      let count = visits ? parseInt(visits, 10) : 0;
      count++;
      await env.SETTINGS?.put(key, count.toString());
      return Response.json({ visits: count });
    }

    // 404 for any other route
    return new Response("Not Found", { status: 404 });
  },
};
