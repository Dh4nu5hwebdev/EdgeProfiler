# EdgeProfiler(v0.1)

**EdgeProfiler is a lightweight HTTP request profiler written in Go. It helps observe incoming HTTP requests and measure how long each request takes to be processed.**

## Why this project exists?

**This project was built to understand how HTTP servers work internally, how request–response cycles are handled, and how basic profiling and instrumentation can be implemented using only Go’s standard library.**

## How it works (high level)

The server listens on port 8080 and waits for incoming requests.
Each request is handled by a handler function where the request start time is recorded.

After processing the request, the elapsed time is calculated and logged along with request metadata.

## How to run?

1. Ensure Go is installed in your device.
2. Navigate to project directory (Edge Profiler)
3. Run the command:

```bash
go run main.go
```

 Open a browser and visit [http://localhost:8080](http://localhost:8080)

### Example Output

```bash
GET / took 10ms
GET /test took 8ms

```

## Current version

v0.1 – Basic request profiling and logging

### Planned improvements

- Forward requests to a backend server (proxy mode)
- Measure backend API latency
- Improve log formatting
- Add request identifiers
