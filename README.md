# Documentation Agent

A Context7 A2A Agent for Inference Gateway built using the Agent Definition Language (ADL).

## Overview

This agent provides documentation search capabilities through the Context7 MCP server integration, enabling developers to access always up-to-date documentation when executing tasks.

## Quick Start

### Local Development

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the agent:**
   ```bash
   go run main.go
   ```

3. **Test the agent:**
   ```bash
   # Health check
   curl http://localhost:8080/health
   
   # List available tools
   curl http://localhost:8080/tools
   ```

### Using with Inference Gateway

```bash
infer a2a connect oci://ghcr.io/inference-gateway/context7-agent:latest
```

## Configuration

The agent is configured via `agent.yaml` using the Agent Definition Language (ADL) specification:

- **Kind**: Agent (ADL v1)
- **Name**: documentation-agent
- **Version**: 0.1.0
- **Protocol**: A2A with HTTP/SSE async transport
- **Port**: 8080

## Available Tools

- **resolve_library_id**: Resolves library information by ID
- **get_library_docs**: Retrieves documentation for a specific library

## Deployment

### Docker

```bash
docker build -t context7-agent .
docker run -p 8080:8080 context7-agent
```

### Kubernetes

Deploy using the Inference Gateway Operator with Agent CRD.

## Architecture

Built with:
- **Language**: Go 1.24
- **Protocol**: A2A (Agent-to-Agent)
- **Transport**: HTTP/SSE with streaming support
- **Integration**: Context7 MCP Server

## Development

This project was generated using the ADL CLI:

```bash
adl init agent.yaml
```

## License

This project is part of the Inference Gateway ecosystem.