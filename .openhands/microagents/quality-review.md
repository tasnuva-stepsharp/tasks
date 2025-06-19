---
trigger: ["quality", "review", "llama", "code review"]
---

# Code Quality Check

Use LLaMA 4 locally via OpenHands to:
- Check for anti-patterns, missing error handling, poor abstractions.
- Rate code from 1–10.
- Suggest exact improvements using actionable comments.

Prompt:
```text
Review this code for Go idioms, error handling, abstraction, and testability.


---

### `runtime-check.md`

```md
---
trigger: ["test", "run", "runtime", "debug"]
---

# Runtime Verification

- Run the entire project using:
  ```bash
  go run main.go


---

### `idea-dedup.md`

```md
---
trigger: ["idea", "feature", "init"]
---

# Idea Coordination Rules

Before working on a new feature:
1. Log the feature idea in `.openhands/idea-log.md`.
2. Cross-check with existing entries.
3. If a similar idea exists, improve or extend it — do not duplicate.
4. Use this prompt to confirm:
```text
Is anyone else working on this idea or a related feature?
