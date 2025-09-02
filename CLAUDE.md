# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the **Context7 A2A Agent** for Inference Gateway - an Agent-to-Agent integration that enables the Inference Gateway to interface with Context7 through the A2A protocol. The agent is built using the Agent Definition Language (ADL) and provides documentation search capabilities.

## Architecture

### Core Components

- **agent.yaml**: ADL configuration defining the agent specification, capabilities, tools, and deployment settings
- **PRD.md**: Product Requirements Document detailing the complete technical specification and acceptance criteria

### Agent Definition

The agent is defined using the Agent Definition Language (ADL) with the following key characteristics:

- **Kind**: Agent (ADL v1)
- **Name**: documentation-agent  
- **Version**: 0.1.0
- **Go Module**: github.com/inference-gateway/documentation-agent
- **Protocol**: A2A (Agent-to-Agent) with HTTP/SSE async transport
- **Port**: 8080 (default server port)

### Available Tools

The agent exposes two primary tools for documentation management:

1. **resolve_library_id**: Resolves library information by ID
2. **get_library_docs**: Retrieves documentation for a specific library

## Development Commands

Since this is an ADL-based project, development typically involves working with the Agent Definition CLI:

```bash
# Connect to the agent (when deployed)
infer a2a connect oci://ghcr.io/inference-gateway/context7-agent:latest

# Local testing would use the ADL CLI commands
# (specific commands depend on ADL CLI installation)
```

## Configuration

The agent configuration supports:

- **Streaming**: Enabled
- **State Transition History**: Enabled  
- **Max Tokens**: 4096
- **Temperature**: 0.7
- **Go Version**: 1.24

Key configuration sections in agent.yaml:
- `spec.capabilities`: Defines agent capabilities
- `spec.agent`: LLM and system prompt configuration
- `spec.tools`: Tool definitions with JSON schemas
- `spec.server`: Server configuration (port, debug settings)
- `spec.language.go`: Go-specific settings

## System Prompt

The agent uses this system prompt:
"You are a helpful assistant for managing and searching through Documentations queries. You can use MCP Client context7 for searching docs."

## Deployment

The agent is designed to be deployed:
- **Locally**: Via CLI for development and testing
- **Kubernetes**: Via the Inference Gateway Operator using Agent CRD
- **Container**: As an OCI image at ghcr.io/inference-gateway/context7-agent

## Issue Templates

The project includes standardized GitHub issue templates:
- Bug reports with reproduction steps
- Feature requests with acceptance criteria  
- Refactor requests