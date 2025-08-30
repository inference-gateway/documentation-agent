# PRD: Context7 A2A Agent for Inference Gateway

## 1. Overview

The **Context7 A2A Agent** is an Agent-to-Agent (A2A) integration that enables Inference Gateway to interface directly with the **Context7 MCP server** (the most widely used MCP server today).
It will leverage the **Agent Definition CLI** to declaratively define, test, and deploy the agent.

This agent will allow developers and operators to easily plug Context7 capabilities into multi-agent workflows without writing custom glue code.

---

## 2. Goals

* ✅ Provide **first-class support for Context7** inside the Inference Gateway ecosystem.
* ✅ Allow **declarative agent definitions** (via the Agent Definition CLI) to expose Context7 tools and capabilities.
* ✅ Optimize for **low token usage and efficient context passing**.
* ✅ Ship an **example agent definition** that users can immediately run with `infer agent run`.

---

## 3. Non-Goals

* ❌ Re-implementing Context7 itself.
* ❌ Vendor-specific hacks outside the A2A standard.
* ❌ Supporting STDIO protocol (bad practice).

---

## 4. User Stories

### 4.1 Developer

* *As a developer*, I want to run `infer agent add context7` so I can immediately use Context7 tools in my multi-agent workflows.
* *As a developer*, I want to declaratively define the Context7 agent in `AGENTS.md` or via the CLI, so I don’t have to manually wire up MCP tool calls.
* *As a developer*, I want to chain Context7 with other agents (e.g., Browser, Docs, GitHub) via A2A so I can build compound automations.

### 4.2 Operator

* *As an operator*, I want to run the Context7 A2A Agent in Kubernetes (via the Operator) so it scales automatically and integrates into GitOps.
* *As an operator*, I want metrics on token usage and tool call frequency for cost optimization.

---

## 5. Functional Requirements

### 5.1 Agent Definition

* The agent must be **declared via Agent Definition CLI**:

  ```bash
  infer agent define context7 \
    --server mcp://context7 \
    --description "Access Context7 MCP tools via Inference Gateway" \
    --tools search
  ```
* Must generate an **AGENTS.md entry** (open-standard format) with the tools and their schemas.

### 5.2 Protocol Support

* The agent must fully implement the **A2A protocol** handshake.

### 5.3 Context Handling

* Agent must support **context injection** and retrieval from Context7.
* Support **token optimization strategies** (e.g., summarization of large Context7 responses).

### 5.4 CLI Integration

* Must integrate with `infer agent run context7`.
* Must allow local testing via the CLI before deployment to Kubernetes.

### 5.5 Deployment

* Operator should reconcile `Agent` CRD with rollout annotations on config change.

---

## 6. Technical Design

* **Language**: Go (same as Inference Gateway).
* **Interface**: A2A server exposing skills like "search" the documentation.
* **Transport**: HTTP/SSE async with polling and real-time sync with streaming.
* **Config**:

  ```yaml
apiVersion: core.inference-gateway.com/v1alpha1
kind: Agent
metadata:
  name: context7-agent
  namespace: agents
spec:
  image: ghcr.io/inference-gateway/context7-agent:latest
  timezone: "UTC"
  port: 8080
  host: "0.0.0.0"
  readTimeout: "30s"
  writeTimeout: "30s"
  idleTimeout: "60s"
  logging:
    level: "info"
    format: "json"
  telemetry:
    enabled: true
    metrics:
      enabled: true
      port: 9090
  queue:
    enabled: true
    maxSize: 1000
    cleanupInterval: "5m"
  tls:
    enabled: false
    secretRef: ""
  agent:
    enabled: true
    tls:
      enabled: true
      secretRef: ""
    maxConversationHistory: 10
    maxChatCompletionIterations: 5
    maxRetries: 3
    apiKey:
      secretRef: "your-api-key"
    llm:
      model: "openai/gpt-3.5-turbo"
      maxTokens: 4096
      temperature: "0.7"
      customHeaders:
        - name: "User-Agent"
          value: "Context7 Agent"
      systemPrompt: "You are a helpful assistant for managing and searching through Documentations queries. You can use MCP Client context7 for searching docs."
  env:
    - name: DEMO_MODE
      valueFrom:
        configMapKeyRef:
          name: context7-config
          key: DEMO_MODE
    - name: A2A_AGENT_URL
      valueFrom:
        configMapKeyRef:
          name: context7-config
          key: A2A_AGENT_URL
  ```
* **Error Handling**: Retry failed MCP calls with exponential backoff.
* **Observability**: Export OpenTelemetry traces for tool calls.

---

## 7. Acceptance Criteria

* [ ] Agent can be defined via CLI and generates a valid `AGENTS.md` entry.
* [ ] `infer agent run context7` successfully connects to a running Context7 A2A server.
* [ ] Tools `search` is a skill associated with this agent.
* [ ] Works in **local mode** (CLI) and **Kubernetes mode** (Operator).
* [ ] Exposes metrics (tool calls, latency, token usage).
* [ ] Support for TLS
* [ ] Built using the Agent Definition Language with the ADL CLI
