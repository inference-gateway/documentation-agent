<div align="center">

# Documentation-Agent
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![A2A Protocol](https://img.shields.io/badge/A2A-Protocol-blue?style=flat)](https://github.com/inference-gateway/a2a)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Assistant for managing and searching through Documentations queries**

A production-ready [Agent-to-Agent (A2A)](https://github.com/inference-gateway/a2a) server that provides AI-powered capabilities through a standardized protocol.

</div>

## Quick Start

```bash
# Run the agent
go run .

# Or with Docker
docker build -t documentation-agent .
docker run -p 8080:8080 documentation-agent
```

## Features

- ✅ A2A protocol compliant
- ✅ AI-powered capabilities
- ✅ Streaming support
- ✅ State transition history
- ✅ Production ready
- ✅ Minimal dependencies

## Endpoints

- `GET /.well-known/agent.json` - Agent metadata and capabilities
- `GET /health` - Health check endpoint
- `POST /a2a` - A2A protocol endpoint

## Available Tools
- **resolve_library_id** - Resolves library by its id
- **get_library_docs** - Get the docs for the specific library

## Configuration

Configure the agent via environment variables:

### Core Application Settings

- `ENVIRONMENT` - Deployment environment

### A2A Agent Configuration

#### Server Configuration

- `A2A_SERVER_PORT` - Server port (default: `8080`)
- `A2A_SERVER_READ_TIMEOUT` - Maximum duration for reading requests (default: `120s`)
- `A2A_SERVER_WRITE_TIMEOUT` - Maximum duration for writing responses (default: `120s`)
- `A2A_SERVER_IDLE_TIMEOUT` - Maximum time to wait for next request (default: `120s`)
- `A2A_SERVER_DISABLE_HEALTHCHECK_LOG` - Disable logging for health check requests (default: `true`)

#### LLM Client Configuration

- `A2A_AGENT_CLIENT_PROVIDER` - LLM provider: `openai`, `anthropic`, `groq`, `ollama`, `deepseek`, `cohere`, `cloudflare`
- `A2A_AGENT_CLIENT_MODEL` - Model to use
- `A2A_AGENT_CLIENT_API_KEY` - API key for LLM provider
- `A2A_AGENT_CLIENT_BASE_URL` - Custom LLM API endpoint
- `A2A_AGENT_CLIENT_TIMEOUT` - Timeout for LLM requests (default: `30s`)
- `A2A_AGENT_CLIENT_MAX_RETRIES` - Maximum retries for LLM requests (default: `3`)
- `A2A_AGENT_CLIENT_MAX_TOKENS` - Maximum tokens for LLM responses (default: `4096`)
- `A2A_AGENT_CLIENT_TEMPERATURE` - Controls randomness of LLM output (default: `0.7`)

## Development

```bash
# Generate code from ADL
task generate

# Run tests
task test

# Build the application
task build

# Run linter
task lint

# Format code
task fmt
```

## Deployment

### Docker

```bash
docker build -t documentation-agent:latest .
docker run -p 8080:8080 \
  -e A2A_AGENT_CLIENT_PROVIDER=openai \
  -e A2A_AGENT_CLIENT_API_KEY=your-api-key \
  documentation-agent:latest
```

### Kubernetes

```bash
kubectl apply -f k8s/
```

## License

MIT License - see LICENSE file for details
