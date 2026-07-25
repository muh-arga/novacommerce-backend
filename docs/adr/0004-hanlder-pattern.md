# ADR-0004: Handler Pattern

## Context
HTTP endpoints will continue growing.

Inline route handlers will become difficult to maintain.

## Decision
Each domain owns its own Handler.

Example:

- HealthHandler
- AuthHandler
- ProductHandler
- OrderHandler

Handlers expose methods instead of standalone functions.

Example:

HealthHandler.Index()

instead of

Index()