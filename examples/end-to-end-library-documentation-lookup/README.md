# End-to-end library documentation lookup

Give the agent a bare library name and a question, and let the
**`library-documentation-lookup`** skill drive the whole flow: resolve the ID,
then fetch the docs, in a single turn. This is the recommended path for most
callers - you don't have to know any Context7 IDs up front.

## What it demonstrates

- The **`library-documentation-lookup`** skill orchestrating both tools
  (`resolve_library_id` → `get_library_docs`) end to end.
- How skills (system-prompt playbooks) and tools (function calls) work
  together.

## Prerequisites

```bash
export A2A_AGENT_CLIENT_PROVIDER=openai
export A2A_AGENT_CLIENT_MODEL=gpt-4o-mini
export A2A_AGENT_CLIENT_API_KEY=sk-...
export CONTEXT7_API_KEY=ctx7_...   # optional, recommended

task run   # or: go run . start
```

Confirm the skill is loaded:

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 tasks submit "What are your skills?"
```

## Run it

```bash
docker run --rm -it --network host \
  ghcr.io/inference-gateway/a2a-debugger:latest \
  --server-url http://localhost:8080 \
  tasks submit "How do I stream responses from the OpenAI Python SDK?"
```

## What happens under the hood

1. The skill recognises a documentation request and reads its own `SKILL.md`
   playbook (via the built-in `Read` tool).
2. `resolve_library_id` maps `"OpenAI Python SDK"` → e.g. `/openai/openai-python`.
3. `get_library_docs` fetches streaming-specific docs for that ID.
4. The agent answers with a worked example grounded in the fetched docs.

All four steps happen inside one task - the caller supplied only a name and a
question.

## Related examples

- [Resolve a library name to a Context7 ID](../resolve-a-library-name-to-a-context7-id/) - step 2 in isolation.
- [Fetch topic-scoped documentation for a library](../fetch-topic-scoped-documentation-for-a-library/) - the two-tool chain, run manually.
- [Look up version-specific API behavior](../look-up-version-specific-api-behavior/) - skip resolution with a known ID.
