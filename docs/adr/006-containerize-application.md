# ADR-0006: Containerize the Go Application

## Context

Developers may use different Go versions and local environments, making application setup inconsistent.

The application needs a reproducible runtime that can be started without requiring a locally installed Go toolchain.

## Decision

The NovaCommerce API will be packaged as a Docker image using a multi-stage Dockerfile.

The build stage compiles the Go application, while the runtime stage contains only Alpine Linux and the compiled binary.

Docker Compose will be used to build and run the API during local development.

The runtime container must:

- Run as a non-root user
- Receive application configuration through environment variables
- Avoid embedding `.env` or other secrets in the image
- Expose the HTTP application port