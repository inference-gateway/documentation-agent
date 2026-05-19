---
name: library-documentation-lookup
description: Use this when you need up-to-date documentation for a third-party library or framework before writing code against it. First resolves the library name to a Context7-compatible ID via resolve_library_id when the caller does not already know it (format '/org/project' or '/org/project/version'), then fetches focused, topic-scoped documentation via get_library_docs. Good for filling in unknowns about specific APIs, hooks, configuration options, or version-specific behavior.
tags:
  - docs
  - libraries
  - context7
  - documentation
  - reference
---

# library-documentation-lookup

Use this when the caller needs current, authoritative documentation for a
third-party library or framework - typically before generating code, picking
an API, or answering a "how do I use X in Y" question. The skill chains
`resolve_library_id` (name → Context7 ID) and `get_library_docs` (ID → docs
text) so the model never has to invent an identifier.

## When to use

Trigger this skill when the request matches any of:

- The caller names a library/framework/SDK and asks how to use it, configure
  it, call a specific API, or migrate between versions
  (e.g. "show me Next.js app-router routing", "how do Supabase RLS policies
  work", "Mongo aggregation pipeline examples").
- The caller pastes a Context7-compatible ID (`/org/project` or
  `/org/project/version`) and asks for docs on it.
- The caller is writing code that imports a package and you need ground
  truth on a function signature, hook, option, or env var before answering.
- The caller asks for docs scoped to a topic ("hooks", "routing",
  "authentication", "streaming") - pass it through as `topic`.

Skip this skill when:

- The question is conceptual and version-agnostic and you already have
  high-confidence knowledge (basic language syntax, well-known algorithms).
- The caller is asking about *this* agent's own behavior, not a third-party
  library.

## Workflow

1. **Identify the library and intent.** Extract the library/framework name
   from the request and form a concrete `query` describing what the caller
   needs to learn (e.g. "How to set up JWT auth in Express.js", "React
   useEffect cleanup examples"). Specific queries return better docs than
   single keywords. If the caller already provided a Context7 ID
   (`/org/project[/version]`), skip step 2.
2. **Resolve the ID.** Call `resolve_library_id` with `libraryName` (use
   official punctuation: `Next.js`, not `nextjs`) **and** `query`. The
   response returns `selectedLibraryID` (best match) plus `allMatches`. If
   `allMatches` contains several plausible candidates and the caller's
   intent is ambiguous (e.g. "react" could be `/facebook/react` or
   `/reactjs/react.dev`), surface the top matches and ask the caller to
   pick before fetching docs.
3. **Fetch the docs.** Call `get_library_docs` with:
   - `libraryId`: the resolved (or caller-supplied) ID.
   - `query`: the same specific question from step 1. Be specific - good:
     "useEffect cleanup function examples"; bad: "hooks".
4. **Answer from the returned `documentation` field.** Quote or summarise
   the relevant portion. Include the resolved `libraryID` so the caller
   can verify which library and version you used.

Do not call either tool more than 3 times for a single request - if you
can't find what you need after 3 attempts, summarise what you have and
ask the caller to refine.

## Tools

Declared under `spec.tools` in `agent.yaml`:

1. `resolve_library_id` - (`libraryName`, `query`) → Context7-compatible
   ID. Always call first unless the caller already supplied an ID.
2. `get_library_docs` - (`libraryId`, `query`) → documentation text.

## Error handling

- **`resolve_library_id` returns no matches:** ask the caller to confirm
  the spelling, suggest the closest match if any, or request the
  Context7 ID directly. Do not guess an ID.
- **`get_library_docs` returns `{"error": "Library not found: ..."}`:**
  the ID is malformed or unpublished. Re-run `resolve_library_id`, or ask
  the caller for the correct ID.
- **`get_library_docs` returns `{"error": "Invalid Context7 API key..."}`:**
  the server is missing `CONTEXT7_API_KEY`. Report this back to the caller
  - it is an operator configuration issue, not something to retry.
- **Either tool returns a `"note": "Using mock response - MCP integration
  pending"`:** the upstream MCP server did not return a `result` field.
  Treat the payload as low-confidence and tell the caller the response is
  a fallback.
