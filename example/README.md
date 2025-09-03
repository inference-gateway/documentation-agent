# Documentation Agent Example

## Overview

This example demonstrates how to run the Context7 Documentation Agent locally using Docker Compose. The setup includes:

- **Documentation Agent**: The A2A agent that provides documentation search capabilities
- **Inference Gateway**: Acts as the inference service provider for the agent
- **A2A Debugger**: A CLI tool for testing and debugging agent interactions

The Documentation Agent uses the Agent-to-Agent (A2A) protocol to interface with Context7, enabling documentation search and retrieval capabilities through a standardized API.

## Prerequisites

- Docker and Docker Compose installed
- API keys for your chosen inference provider (e.g., DeepSeek)
- Basic understanding of the A2A protocol

## Quick Start

### 1. Configuration

Copy the example environment files and configure them with your API keys:

```bash
cp .env.agent.example .env.agent
cp .env.gateway.example .env.gateway
```

#### Environment Variables

**.env.agent** - Agent configuration:
- `ENVIRONMENT`: Set to `development` or `production`
- `CONTEXT7_API_KEY`: Your Context7 API key for documentation access
- `A2A_AGENT_URL`: Agent server URL (default: `http://localhost:8080`)
- `A2A_AGENT_CLIENT_PROVIDER`: LLM provider (e.g., `deepseek`)
- `A2A_AGENT_CLIENT_MODEL`: Model to use (e.g., `deepseek/deepseek-chat`)
- `A2A_AGENT_CLIENT_BASE_URL`: Inference service URL (default: `http://inference-gateway:8080/v1`)

**.env.gateway** - Inference Gateway configuration:
- `ENVIRONMENT`: Set to `development` or `production`
- `DEEPSEEK_API_KEY`: Your DeepSeek API key (or other provider's key)

### 2. Start the Services

Launch all services with Docker Compose:

```bash
docker compose up --build
```

This will start:
- Documentation Agent on port 8080
- Inference Gateway (internal network)
- Services will auto-restart unless stopped

### 3. Test the Agent

Submit a test query using the A2A debugger:

```bash
docker compose run --rm a2a-debugger tasks submit-streaming "What's the latest version of NextJS?"
```

## Usage Examples

### Basic Documentation Query
```bash
docker compose run --rm a2a-debugger tasks submit-streaming "How do I use React hooks?"
```

### Library-Specific Search
```bash
docker compose run --rm a2a-debugger tasks submit-streaming "Search Vue.js documentation for composition API"
```

### Submit Non-Streaming Task
```bash
docker compose run --rm a2a-debugger tasks submit "What is TypeScript?"
```

## Architecture

```
┌─────────────┐     A2A Protocol    ┌──────────────────┐
│ A2A         │◄───────────────────►│ Documentation    │
│ Debugger    │                     │ Agent            │
└─────────────┘                     └──────────────────┘
                                              │
                                              │ HTTP/REST
                                              ▼
                                     ┌──────────────────┐
                                     │ Inference        │
                                     │ Gateway          │
                                     └──────────────────┘
                                              │
                                              ▼
                                     ┌──────────────────┐
                                     │ LLM Provider     │
                                     │ (DeepSeek, etc)  │
                                     └──────────────────┘
```

## Available Tools

The Documentation Agent exposes two primary tools:

1. **resolve_library_id**: Resolves library information by ID
2. **get_library_docs**: Retrieves documentation for a specific library

View all available tools:
```bash
docker compose run --rm a2a-debugger tools list
```

## Troubleshooting

### Agent Not Responding
- Check if all services are running: `docker compose ps`
- View agent logs: `docker compose logs agent`
- Ensure API keys are correctly set in `.env` files

### API Key Issues
- Ensure `CONTEXT7_API_KEY` is valid and has proper permissions
- Verify your inference provider API key (e.g., `DEEPSEEK_API_KEY`) is active

### Debugging Tips
- Enable debug mode by setting `DEBUG=true` in `.env.agent`
- Monitor real-time logs: `docker compose logs -f`
- Test individual services: `docker compose up agent` (run only the agent)

## Development

### Building from Source
```bash
# Build only the agent
docker compose build agent

# Build with no cache
docker compose build --no-cache
```

### Using a Different Inference Provider
You can bypass the Inference Gateway and connect directly to any OpenAI-compatible API by modifying `A2A_AGENT_CLIENT_BASE_URL` in `.env.agent`:

```bash
# Example: Direct OpenAI connection
A2A_AGENT_CLIENT_BASE_URL=https://api.openai.com/v1
A2A_AGENT_CLIENT_PROVIDER=openai
A2A_AGENT_CLIENT_MODEL=gpt-4
```

### Running Services Individually
```bash
# Start only the agent
docker compose up agent

# Start only the gateway
docker compose up inference-gateway

# Run debugger commands manually
docker compose run --rm a2a-debugger --help
```

## Cleanup

Stop all services:
```bash
docker compose down
```

Remove volumes and networks:
```bash
docker compose down -v
```

## Additional Resources

- [A2A Protocol Documentation](https://github.com/inference-gateway/a2a-protocol)
- [Agent Definition Language (ADL) Spec](https://github.com/inference-gateway/adl)
- [Inference Gateway Documentation](https://github.com/inference-gateway/inference-gateway)
- [Context7 API Documentation](https://docs.context7.com)

## License

See the main repository LICENSE file for details.
