# AGENTS.md

This file describes the agents available in this A2A (Agent-to-Agent) system.

## Agent Overview

### documentation-agent
**Version**: 0.2.23  
**Description**: A2A agent server that provides Context7-style documentation capabilities for your agents

This agent is built using the Agent Definition Language (ADL) and provides A2A communication capabilities.

## Agent Capabilities



- **Streaming**: ✅ Real-time response streaming supported


- **Push Notifications**: ❌ Server-sent events not supported


- **State History**: ✅ Tracks agent state transitions over time



## AI Configuration





**System Prompt**: You are an intelligent documentation retrieval assistant that specializes in finding and fetching relevant documentation from Context7-compatible sources. You can resolve library names to their proper identifiers and retrieve targeted documentation based on specific topics or requirements.



**Configuration:**

- Max Tokens: 4096


- Temperature: 0.7



## Skills


This agent provides 2 skills:


### resolve_library_id
- **Description**: Resolves library name to Context7-compatible library ID and returns matching libraries
- **Tags**: docs, libraries
- **Input Schema**: Defined in agent configuration
- **Output Schema**: Defined in agent configuration


### get_library_docs
- **Description**: Fetches up-to-date documentation for a library using Context7-compatible library ID
- **Tags**: docs, libraries
- **Input Schema**: Defined in agent configuration
- **Output Schema**: Defined in agent configuration




## Server Configuration

**Port**: 8080

**Debug Mode**: ❌ Disabled



**Authentication**: ❌ Not required


## API Endpoints

The agent exposes the following HTTP endpoints:

- `GET /.well-known/agent-card.json` - Agent metadata and capabilities
- `POST /skills/{skill_name}` - Execute a specific skill
- `GET /skills/{skill_name}/stream` - Stream skill execution results
- `GET /history` - Retrieve agent state transition history

## Environment Setup

### Required Environment Variables

Key environment variables you'll need to configure:



- `PORT` - Server port (default: 8080)

### Development Environment


**Flox Environment**: ✅ Configured for reproducible development setup




## Usage

### Starting the Agent

```bash
# Install dependencies
go mod download

# Run the agent
go run main.go

# Or use Task
task run
```


### Communicating with the Agent

The agent implements the A2A protocol and can be communicated with via HTTP requests:

```bash
# Get agent information
curl http://localhost:8080/.well-known/agent-card.json



# Execute resolve_library_id skill
curl -X POST http://localhost:8080/skills/resolve_library_id \
  -H "Content-Type: application/json" \
  -d '{"input": "your_input_here"}'

# Execute get_library_docs skill
curl -X POST http://localhost:8080/skills/get_library_docs \
  -H "Content-Type: application/json" \
  -d '{"input": "your_input_here"}'


```

## Deployment


**Deployment Type**: Manual
- Build and run the agent binary directly
- Use provided Dockerfile for containerized deployment



### Docker Deployment
```bash
# Build image
docker build -t documentation-agent .

# Run container
docker run -p 8080:8080 documentation-agent
```


## Development

### Project Structure

```
.
├── main.go              # Server entry point
├── skills/              # Business logic skills

│   └── resolve_library_id.go   # Resolves library name to Context7-compatible library ID and returns matching libraries

│   └── get_library_docs.go   # Fetches up-to-date documentation for a library using Context7-compatible library ID

├── .well-known/         # Agent configuration
│   └── agent-card.json  # Agent metadata
├── go.mod               # Go module definition
└── README.md            # Project documentation
```


### Testing

```bash
# Run tests
task test
go test ./...

# Run with coverage
task test:coverage
```


## Contributing

1. Implement business logic in skill files (replace TODO placeholders)
2. Add comprehensive tests for new functionality
3. Follow the established code patterns and conventions
4. Ensure proper error handling throughout
5. Update documentation as needed

## Agent Metadata

This agent was generated using ADL CLI v0.2.23 with the following configuration:

- **Language**: Go
- **Template**: Minimal A2A Agent
- **ADL Version**: adl.dev/v1

---

For more information about A2A agents and the ADL specification, visit the [ADL CLI documentation](https://github.com/inference-gateway/adl-cli).
