# Configuration

The agent is configured through environment variables (most use the `A2A_`
prefix). This page lists the settings most relevant to
`documentation-agent`; see the project README for the complete table.

## LLM provider

The agent uses an OpenAI-compatible client to drive tool calls.

| Variable | Description | Example |
|----------|-------------|---------|
| `A2A_AGENT_CLIENT_PROVIDER` | LLM provider: `openai`, `anthropic`, `azure`, `ollama`, `deepseek` | `deepseek` |
| `A2A_AGENT_CLIENT_MODEL` | Model identifier | `deepseek-v4-flash` |
| `A2A_AGENT_CLIENT_API_KEY` | Provider API key | `sk-...` |
| `A2A_AGENT_CLIENT_BASE_URL` | Custom endpoint (e.g. an Inference Gateway) | `http://inference-gateway:8080/v1` |

## Context7 access

The `resolve_library_id` and `get_library_docs` tools call Context7. Provide
an API key for authenticated access:

| Variable | Description |
|----------|-------------|
| `CONTEXT7_API_KEY` | Context7 API key. Optional — the agent logs a warning and proceeds unauthenticated when unset, but requests may be rate-limited or rejected with HTTP 401. |

## Server

| Variable | Description | Default |
|----------|-------------|---------|
| `A2A_PORT` | Server port | `8080` |
| `A2A_DEBUG` | Enable debug logging | `false` |
| `A2A_AGENT_URL` | Public URL used in the agent card | `http://localhost:8080` |
| `A2A_SERVER_READ_TIMEOUT` / `A2A_SERVER_WRITE_TIMEOUT` / `A2A_SERVER_IDLE_TIMEOUT` | HTTP timeouts | `120s` |

## Capabilities

| Variable | Description | Default |
|----------|-------------|---------|
| `A2A_CAPABILITIES_STREAMING` | Stream responses over SSE | `true` |
| `A2A_CAPABILITIES_STATE_TRANSITION_HISTORY` | Record task state transitions | `true` |

A ready-to-copy starting point lives in
[`example/.env.agent.example`](../example/.env.agent.example).
