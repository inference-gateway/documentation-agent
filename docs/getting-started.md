# Getting Started

`documentation-agent` is an A2A (Agent-to-Agent) server that provides
Context7-style documentation retrieval. It resolves library names to
Context7-compatible IDs and fetches up-to-date, topic-scoped documentation
that other agents can consume over the A2A protocol.

## Prerequisites

- Go 1.26.4+ (or Docker)
- An OpenAI-compatible LLM endpoint (see [Configuration](configuration.md))
- A [Context7](https://context7.com) API key (optional but recommended)

## Install and run

```bash
# Run directly
go run . start

# Or build the CLI and run the binary
task build
./bin/documentation-agent start

# Or with Docker
docker build -t documentation-agent .
docker run -p 8080:8080 documentation-agent
```

The server listens on port `8080` by default (override with `A2A_PORT`).

## Configure the LLM provider

The agent needs an LLM to decide when to call its tools. At minimum set:

```bash
export A2A_AGENT_CLIENT_PROVIDER=deepseek      # openai | anthropic | azure | ollama | deepseek
export A2A_AGENT_CLIENT_MODEL=deepseek-v4-flash
export A2A_AGENT_CLIENT_API_KEY=<your-key>     # or A2A_AGENT_CLIENT_BASE_URL for a gateway
```

See [Configuration](configuration.md) for the full list.

## Send your first query

With the server running, use the A2A Debugger to submit a task:

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 \
  tasks submit "What is the Context7 ID for Next.js?"
```

Check that the agent is healthy and inspect its capabilities:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/.well-known/agent-card.json
```

Next: [Configuration](configuration.md) · [Usage](usage.md)
