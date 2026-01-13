# Project Rules

> These rules are **non-negotiable**. All AI agents and human contributors must follow them.

---

## Scope Control

1. **MVP only.** Do not add features beyond the defined scope in `TASKS.md`.
2. **No authentication.** Users are anonymous. No login, no registration, no sessions.
3. **No persistence of game results.** Game sessions exist only in memory.
4. **No teams, chat, or social features.**
5. **Single deployment target.** One server, one database, one frontend build.

---

## Architecture Rules

6. **Monolith.** No microservices. No service mesh. No message queues.
7. **No Redis.** In-memory state is sufficient for MVP scale.
8. **No ORM.** Use MongoDB driver directly.
9. **No GraphQL.** REST for CRUD, WebSocket for real-time.
10. **Frontend routing only.** No server-side rendering.

---

## Code Rules

11. **No abstractions without immediate use.** Do not create interfaces, factories, or patterns "for future flexibility."
12. **No dead code.** If it's not used, delete it.
13. **No TODO comments.** Either do it now or add it to `TASKS.md`.
14. **Explicit over clever.** Prefer 10 clear lines over 3 cryptic ones.
15. **English only.** All code, comments, and docs in English.

---

## AI Agent Rules

16. **Read `DOMAIN.md` before generating code.** Use canonical terms only.
17. **Read `ARCHITECTURE.md` before suggesting changes.** Respect system boundaries.
18. **Do not introduce new dependencies** without explicit approval.
19. **Do not refactor working code** unless explicitly requested.
20. **Small changes.** One logical change per commit. No "cleanup while I'm here."

---

## Violation Handling

If any rule is violated:
1. Stop.
2. Revert the violation.
3. Ask for clarification.

These rules exist to prevent the project from stalling again.
