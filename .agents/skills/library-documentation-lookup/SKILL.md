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

TODO: Describe when and how the agent should use this skill. Lead with an
action-oriented "Use this when…" sentence so the model can decide whether
to apply it. The full body of this file is prepended to the system prompt
at runtime.

## When to use

Describe the user intents or task shapes that should trigger this skill.
Be concrete - list the kinds of requests, signals, or context that map to
this playbook.

## Workflow

1. ...
2. ...
3. ...

## Tools

List the tools this skill expects to call (declared under `spec.tools` in
the ADL manifest), and the order in which they're typically invoked.

## Bundled assets

This skill lives in its own directory under `.agents/skills/library-documentation-lookup/`
(also reachable as `.claude/skills/library-documentation-lookup/` via the generated
`.claude/skills` -> `../.agents/skills` symlink). You can ship arbitrary scripts, templates, or
reference material alongside `SKILL.md` - the `.adl-ignore` file protects
the whole directory from being clobbered on regeneration. Suggested layout:

```
.agents/skills/library-documentation-lookup/
├── SKILL.md          # this file
├── scripts/          # optional helper scripts (Python, shell, etc.)
├── templates/        # optional file templates the agent can fill in
└── resources/        # optional static reference material
```

Reference bundled files by relative path from `SKILL.md` (e.g.
`scripts/triage.py`, `templates/report.md`) so the agent can locate them
at runtime.
