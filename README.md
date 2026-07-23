<div align="center">

# Documentation Agent

[![CI](https://github.com/inference-gateway/documentation-agent/workflows/CI/badge.svg)](https://github.com/inference-gateway/documentation-agent/actions/workflows/ci.yml)
[![Go Report Card](https://img.shields.io/badge/Go%20Report%20Card-A+-brightgreen?style=flat&logo=go&logoColor=white)](https://goreportcard.com/report/github.com/inference-gateway/documentation-agent)
[![Go Version](https://img.shields.io/badge/Go-1.26.4+-00ADD8?style=flat&logo=go)](https://golang.org)
[![A2A Protocol](https://img.shields.io/badge/A2A-Protocol-blue?style=flat)](https://github.com/inference-gateway/adk)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

**A2A agent server that provides Context7-style documentation capabilities for your agents**

A enterprise-ready [Agent-to-Agent (A2A)](https://github.com/inference-gateway/adk) server that provides AI-powered capabilities through a standardized protocol.

</div>

## Quick Start

The generated binary is a CLI. `start` boots the A2A server; `--help` and
`--version` work as you'd expect.

```bash
# Run the agent
go run . start

# Or build and invoke the CLI directly
task build
./bin/documentation-agent start

# Or with Docker
docker build -t documentation-agent .
docker run -p 8080:8080 documentation-agent
```

### CLI

| Command | Description |
|---------|-------------|
| `documentation-agent start` | Start the A2A server (blocks until SIGINT/SIGTERM) |
| `documentation-agent --help` | Show top-level help (and per-subcommand with `<cmd> --help`) |
| `documentation-agent --version` | Print the embedded version and exit |

## Quick Install

Add this agent to your Inference Gateway CLI:

```bash
infer agents add documentation-agent http://localhost:8080 \
  --oci ghcr.io/inference-gateway/documentation-agent:latest \
  --run
```

## Features

- ✅ A2A protocol compliant
- ✅ AI-powered capabilities
- ✅ Streaming support
- ✅ State transition history
- ✅ OpenTelemetry instrumentation
- ✅ Enterprise-ready
- ✅ Minimal dependencies

## Endpoints

- `GET /.well-known/agent-card.json` - Agent metadata and capabilities
- `GET /health` - Health check endpoint
- `POST /a2a` - A2A protocol endpoint

## Available Tools

| Tool | Description | Parameters |
|------|-------------|------------|
| `Read` | Read a file from disk. Returns its contents, optionally sliced by line offset/limit. Use this to load SKILL.md bodies on demand. | file_path, offset, limit |
| `resolve_library_id` | Resolves library name to Context7-compatible library ID and returns matching libraries | libraryName, query |
| `get_library_docs` | Fetches up-to-date documentation for a library using Context7-compatible library ID | libraryId, query |

## Examples

| Example | Description |
|---------|-------------|
| [Resolve a library name to a Context7 ID](examples/resolve-a-library-name-to-a-context7-id/) | Ask "What is the Context7 ID for Next.js?" and the agent calls resolve_library_id to return matching libraries with their '/org/project' identifiers. |
| [Fetch topic-scoped documentation for a library](examples/fetch-topic-scoped-documentation-for-a-library/) | Ask "How do I set up JWT auth in Express.js?" and the agent resolves the library, then calls get_library_docs to return focused, up-to-date documentation for that specific topic. |
| [Look up version-specific API behavior](examples/look-up-version-specific-api-behavior/) | Provide a versioned ID such as '/vercel/next.js/v14.3.0-canary.87' and ask about the App Router; the agent fetches documentation scoped to that exact version. |
| [End-to-end library documentation lookup](examples/end-to-end-library-documentation-lookup/) | Give a bare library name and a question; the library-documentation-lookup skill guides the agent to resolve the ID first, then retrieve the relevant docs in a single flow. |

## Skills (loaded into the system prompt)

| Skill | Description | Source |
|-------|-------------|--------|
| `library-documentation-lookup` | Use this when you need up-to-date documentation for a third-party library or framework before writing code against it. First resolves the library name to a Context7-compatible ID via resolve_library_id when the caller does not already know it (format '/org/project' or '/org/project/version'), then fetches focused, topic-scoped documentation via get_library_docs. Good for filling in unknowns about specific APIs, hooks, configuration options, or version-specific behavior. | bare scaffold (`.agents/skills/library-documentation-lookup/SKILL.md`) |

## Documentation
- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [Usage](docs/usage.md)

## Configuration

The agent is configured via environment variables. Defaults are derived
from `agent.yaml`; see [CONFIGURATIONS.md](CONFIGURATIONS.md) for the
full reference of custom and `A2A_*` variables.

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

### Adding Dependencies

The generator owns the baseline toolchain pins (SDK, server framework,
logging, CLI, sandbox utilities). To extend the project without forking
the templates, declare extras in `agent.yaml` - every empty list below
is rendered by `adl init --defaults` precisely so it's discoverable:

| Where | Purpose | Example entry | Rendered into |
|-------|---------|---------------|---------------|
| `spec.language.go.vendor.deps` | Runtime Go modules | `github.com/stretchr/testify@v1.10.0` | `go.mod` `require` block |
| `spec.language.go.vendor.devdeps` | Executable dev tools (Go 1.24 `tool` directive) | `golang.org/x/tools/cmd/stringer@v0.20.0` | `go.mod` `tool` directive |
| `spec.development.deps` | Cross-cutting sandbox tools (not tied to one language) | `kubectl@1.31.0`, `terraform@1.9.5`, `deno@2.1.4` | Flox `manifest.toml` / devcontainer feature |

Entries use the `<package>@<version>` form. Built-in pins always win on
conflict; the generator prints a warning and skips the user entry when
shadowing is attempted. After editing `agent.yaml`, re-run `task generate`
to refresh the manifests.

### Debugging

Use the [A2A Debugger](https://github.com/inference-gateway/a2a-debugger) to test and debug your A2A agent during development. It provides a web interface for sending requests to your agent and inspecting responses, making it easier to troubleshoot issues and validate your implementation.

```bash
docker run --rm -it --network host ghcr.io/inference-gateway/a2a-debugger:latest --server-url http://localhost:8080 tasks submit "What are your skills?"
```

```bash
docker run --rm -it --network host ghcr.io/inference-gateway/a2a-debugger:latest --server-url http://localhost:8080 tasks list
```

```bash
docker run --rm -it --network host ghcr.io/inference-gateway/a2a-debugger:latest --server-url http://localhost:8080 tasks get <task ID>
```

## Deployment

### Docker

The Docker image can be built with custom version information using build arguments:

```bash
docker build \
  --build-arg VERSION=1.2.3 \
  --build-arg AGENT_NAME="My Custom Agent" \
  --build-arg AGENT_DESCRIPTION="Custom agent description" \
  -t documentation-agent:1.2.3 .
```

**Available Build Arguments:**

- `VERSION` - Agent version (default: `0.3.0`)
- `AGENT_NAME` - Agent name (default: `documentation-agent`)
- `AGENT_DESCRIPTION` - Agent description (default: `A2A agent server that provides Context7-style documentation capabilities for your agents`)

These values are embedded into the binary at build time using linker flags, making them accessible at runtime without requiring environment variables.

## License

Apache 2.0 License - see LICENSE file for details
