# ADR-0005: HTTP Response

## Context
API responses must remain consistent across all modules.

## Decision
Create a shared response package responsible for formatting all HTTP responses.

Handlers should never build JSON directly.

Only two response helpers will exist:

- Success()
- Error()

HTTP status code is provided by the caller.