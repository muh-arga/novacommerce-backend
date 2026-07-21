# ADR-0001: Project Structure

## Context
The project requires a scalable folder structure that can grow from a simple REST API into a production-ready backend supporting caching, queues, modular monolith, and eventually microservices.

---
## Decision

### Use Package by Feature (Domain-Driven Folder Strcuture)

```
cmd/
configs/
internal/
pkg/
docs/
```

### Use shared for cross-cutting concerns

For cross-cutting concerns will stored on shared (midlleware, etc.).