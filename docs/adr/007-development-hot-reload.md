# ADR-0007: Use Air for Development Hot Reload

## Context

The production-style Docker image compiles the application into a binary.

During active development, every source-code change would require manually rebuilding and restarting the container. This creates a slow feedback loop.

A development workflow is required that can detect Go source changes and restart the application automatically.

## Decision

Use Air as the development hot-reload tool.

A separate `Dockerfile.dev` will be maintained for development. It includes the Go toolchain and Air.

Docker Compose will:

- Use `Dockerfile.dev`
- Mount the project source code into the container
- Store downloaded Go modules in a named volume
- Run Air as the container command