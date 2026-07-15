# Usage

`documentation-agent` answers documentation questions about third-party
libraries and frameworks. Send it a task over the A2A protocol and it
resolves the library, fetches the relevant docs, and answers from them.

## Tools

| Tool | Inputs | Purpose |
|------|--------|---------|
| `resolve_library_id` | `libraryName`, `query` | Resolve a library name (e.g. `Next.js`) to a Context7-compatible ID (`/org/project` or `/org/project/version`) and return matching candidates. |
| `get_library_docs` | `libraryId`, `query` | Fetch focused, topic-scoped documentation for a resolved or caller-supplied ID. |
| `Read` | `file_path`, `offset`, `limit` | Read a file from disk (used to load skill playbooks on demand). |

## Skill: library-documentation-lookup

The bundled `library-documentation-lookup` skill tells the agent how to chain
the two tools:

1. Identify the library and form a specific `query`
   (e.g. "How to set up JWT auth in Express.js", not just "auth").
2. Call `resolve_library_id` to get the Context7 ID — unless the caller
   already supplied one in `/org/project[/version]` form.
3. Call `get_library_docs` with that ID and the same specific query.
4. Answer from the returned documentation, citing the library ID used.

The agent will not call the tools more than three times for a single request.

## Example queries

Assuming the agent is running on `http://localhost:8080`:

```bash
DEBUG="docker run --rm -it --network host ghcr.io/inference-gateway/a2a-debugger:latest --server-url http://localhost:8080"

# Resolve a name to a Context7 ID
$DEBUG tasks submit "What is the Context7 ID for Next.js?"

# Fetch topic-scoped documentation
$DEBUG tasks submit "How do I set up JWT auth in Express.js?"

# Ask about a specific version
$DEBUG tasks submit "Show App Router changes in /vercel/next.js/v14.3.0-canary.87"

# Inspect submitted tasks
$DEBUG tasks list
$DEBUG tasks get <task-id>
```

## Troubleshooting

| Symptom | Cause | Fix |
|---------|-------|-----|
| `Invalid Context7 API key` | `CONTEXT7_API_KEY` missing or wrong | Set a valid key (see [Configuration](configuration.md)). |
| `Library not found: ...` | Malformed or unpublished ID | Re-run `resolve_library_id`, or pass a correct `/org/project` ID. |
| `Using mock response - MCP integration pending` | Upstream returned no result | Treat the response as a low-confidence fallback. |
