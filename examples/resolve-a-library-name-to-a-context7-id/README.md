# Resolve a library name to a Context7 ID

Resolve a human-friendly library name (e.g. `Next.js`) into a
Context7-compatible ID such as `/vercel/next.js`. This is the first step of
almost every documentation lookup: the ID is what `get_library_docs` needs to
fetch focused docs.

## What it demonstrates

- A single call to the **`resolve_library_id`** tool.
- How Context7 ranks candidate libraries and returns the best match alongside
  the alternates.

## Prerequisites

Configure an LLM provider and start the agent:

```bash
# Any OpenAI-compatible provider works (openai, anthropic, deepseek, ollama, ...)
export A2A_AGENT_CLIENT_PROVIDER=openai
export A2A_AGENT_CLIENT_MODEL=gpt-4o-mini
export A2A_AGENT_CLIENT_API_KEY=sk-...

# Optional but recommended - lifts Context7 rate limits.
# https://context7.com/docs/howto/api-keys
export CONTEXT7_API_KEY=ctx7_...

task run   # or: go run . start
```

## Run it

Submit a task with the [A2A Debugger](https://github.com/inference-gateway/a2a-debugger):

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 \
  tasks submit "What is the Context7 ID for Next.js?"
```

Then read the result back:

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 tasks list

docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 tasks get <task-id>
```

## What happens under the hood

The agent calls `resolve_library_id` with:

```json
{ "libraryName": "Next.js", "query": "What is the Context7 ID for Next.js?" }
```

which returns the top match plus the alternates it considered:

```json
{
  "selectedLibraryID": "/vercel/next.js",
  "totalMatches": 5,
  "selectedLibrary": { "id": "/vercel/next.js", "title": "Next.js", "trustScore": 10 },
  "allMatches": [ /* ...ranked candidates... */ ]
}
```

The agent then replies with the resolved ID, e.g.
_"The Context7 ID for Next.js is `/vercel/next.js`."_

## Next step

Feed that ID into [Fetch topic-scoped documentation for a
library](../fetch-topic-scoped-documentation-for-a-library/) to retrieve real
docs.
