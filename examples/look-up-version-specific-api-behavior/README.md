# Look up version-specific API behavior

Documentation drifts between releases. When you already know the exact
Context7 ID - including a version suffix - hand it to the agent directly and it
fetches docs pinned to that version, skipping resolution entirely.

## What it demonstrates

- Calling **`get_library_docs`** with a versioned ID
  (`/org/project/version`), bypassing `resolve_library_id`.
- Version-scoped answers: the same question can return different docs for
  different releases.

## Prerequisites

```bash
export A2A_AGENT_CLIENT_PROVIDER=openai
export A2A_AGENT_CLIENT_MODEL=gpt-4o-mini
export A2A_AGENT_CLIENT_API_KEY=sk-...
export CONTEXT7_API_KEY=ctx7_...   # optional, recommended

task run   # or: go run . start
```

## Run it

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 \
  tasks submit "Using /vercel/next.js/v14.3.0-canary.87, how does the App Router handle route groups?"
```

## What happens under the hood

Because a fully-qualified, versioned ID is already present in the request, the
agent goes straight to `get_library_docs`:

```json
{
  "libraryId": "/vercel/next.js/v14.3.0-canary.87",
  "query": "How does the App Router handle route groups?"
}
```

The returned `documentation` reflects that exact canary build rather than the
latest stable release.

## Notes

- ID format is `/org/project` or `/org/project/version`. The tool rejects any
  ID that doesn't start with `/`.
- Don't know the version string? Resolve first - `resolve_library_id` lists a
  library's available `versions` in its match metadata.
