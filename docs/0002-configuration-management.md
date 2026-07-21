# ADR-0002: Configuration Management

## Context
The application requires configuration for:

- Application
- Database
- Redis
- RabbitMQ

## Decision
A centralized configuration package is introduced.

All configuration is loaded once during application startup using Viper and exposed through a single Config object.