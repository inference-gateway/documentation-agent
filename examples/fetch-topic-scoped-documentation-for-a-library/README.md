# Fetch topic-scoped documentation for a library

Ask a focused question about a library and get back documentation scoped to
that exact topic - not the whole manual. The agent resolves the library first,
then fetches only the relevant snippets.

## What it demonstrates

- A two-tool chain: **`resolve_library_id`** → **`get_library_docs`**.
- Topic scoping: the `query` narrows the docs down to the asked-about API.

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
  tasks submit "How do I set up JWT auth in Express.js?"
```

Read the answer back with `tasks list` then `tasks get <task-id>` (see the
[resolve example](../resolve-a-library-name-to-a-context7-id/) for the exact
commands).

## What happens under the hood

1. `resolve_library_id` maps `"Express.js"` → `/expressjs/express`.
2. `get_library_docs` fetches docs scoped to the topic:

```json
{ "libraryId": "/expressjs/express", "query": "How to set up JWT auth in Express.js" }
```

returning:

```json
{
  "libraryID": "/expressjs/express",
  "query": "How to set up JWT auth in Express.js",
  "documentation": "...middleware examples using jsonwebtoken..."
}
```

The agent synthesises those snippets into a direct, code-backed answer.

## Tips

- Be specific in the question. _"JWT auth in Express.js"_ returns tighter, more
  useful snippets than _"auth"_.
- Already know the ID? Skip resolution and pass it straight through - see
  [Look up version-specific API behavior](../look-up-version-specific-api-behavior/).
