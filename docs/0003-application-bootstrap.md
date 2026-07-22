# ADR-0003: Application Bootstrap

## Context
As the project grows, the application will require multiple infrastructure components such as:

- HTTP Router
- Database
- Redis
- RabbitMQ
- Logger

Keeping initialization logic inside `main.go` would make the application difficult to maintain.

## Decision
Introduce an `Application` struct responsible for owning application dependencies.

Responsibilities:

- Load configuration
- Build dependencies
- Initialize router
- Start HTTP server

`main.go` is only responsible for bootstrapping the application.