# Style Guide

> Conventions for code structure, formatting, and patterns. Follow these to maintain consistency.

---

## Folder Structure

### Backend (`backend/`)

```
backend/
├── cmd/
│   └── main.go              # Entry point only
├── internal/
│   ├── app.go               # Bootstrap, DI, server setup
│   ├── config/
│   │   └── config.go        # Environment loading
│   ├── entity/              # Domain structs (Quiz, Question, etc.)
│   ├── handler/             # HTTP + WebSocket handlers
│   │   ├── quiz.go          # Quiz REST endpoints
│   │   └── ws.go            # WebSocket connection handler
│   ├── service/             # Business logic
│   │   ├── quiz.go          # Quiz CRUD logic
│   │   └── game.go          # Game session management
│   ├── repository/          # Data access (MongoDB)
│   │   └── quiz.go
│   └── ws/                   # WebSocket specifics
│       ├── hub.go           # Connection management
│       ├── message.go       # Message types
│       └── handler.go       # Event routing
├── go.mod
└── go.sum
```

**Rules:**
- `cmd/` contains only `main.go`. No logic.
- `internal/` is private. No exports outside this module.
- `entity/` has no dependencies on other packages.
- `handler/` calls `service/`. Never calls `repository/` directly.
- `service/` calls `repository/`. Contains business rules.
- `repository/` talks to MongoDB. Returns entities.

### Frontend (`frontend/`)

```
frontend/
├── src/
│   ├── main.ts              # Mount app
│   ├── App.svelte           # Root component, routing
│   ├── routes/              # Page components
│   │   ├── Home.svelte
│   │   ├── host/
│   │   │   ├── Dashboard.svelte
│   │   │   ├── Lobby.svelte
│   │   │   └── Game.svelte
│   │   └── play/
│   │       ├── Join.svelte
│   │       ├── Lobby.svelte
│   │       └── Game.svelte
│   ├── components/          # Reusable UI pieces
│   │   ├── Button.svelte
│   │   ├── QuizCard.svelte
│   │   ├── Timer.svelte
│   │   └── Leaderboard.svelte
│   ├── stores/              # Svelte stores (state)
│   │   ├── game.ts
│   │   └── websocket.ts
│   ├── services/            # API/WS clients
│   │   ├── api.ts
│   │   └── socket.ts
│   └── app.css              # Global styles
├── index.html
├── package.json
├── tailwind.config.js
└── vite.config.ts
```

**Rules:**
- `routes/` are full pages. One per URL path.
- `components/` are reusable. No direct API calls.
- `stores/` hold reactive state. No UI logic.
- `services/` handle external communication.

---

## Naming Conventions

### Files

| Type | Convention | Example |
|------|------------|---------|
| Go files | lowercase, short | `quiz.go`, `game.go` |
| Svelte components | PascalCase | `QuizCard.svelte` |
| TypeScript modules | camelCase | `websocket.ts` |

### Go

| Type | Convention | Example |
|------|------------|---------|
| Exported functions | PascalCase | `CreateQuiz` |
| Unexported functions | camelCase | `generateRoomCode` |
| Structs | PascalCase | `GameSession` |
| Constants | PascalCase | `MaxPlayers` |
| Variables | camelCase | `currentQuestion` |

### TypeScript

| Type | Convention | Example |
|------|------------|---------|
| Functions | camelCase | `joinGame` |
| Classes | PascalCase | `WebSocketClient` |
| Interfaces | PascalCase, no prefix | `Player`, not `IPlayer` |
| Constants | UPPER_SNAKE_CASE | `API_BASE_URL` |

---

## Error Handling

### Go

```go
// DO: Return errors explicitly
func GetQuiz(id string) (*Quiz, error) {
    quiz, err := repo.FindByID(id)
    if err != nil {
        return nil, fmt.Errorf("get quiz %s: %w", id, err)
    }
    return quiz, nil
}

// DON'T: Panic or ignore errors
func GetQuiz(id string) *Quiz {
    quiz, _ := repo.FindByID(id)  // BAD
    return quiz
}
```

**Rules:**
- Always wrap errors with context: `fmt.Errorf("action: %w", err)`
- HTTP handlers return appropriate status codes
- Log errors at handler level, not service level

### TypeScript

```typescript
// DO: Handle errors at call site
try {
    const quizzes = await api.getQuizzes();
    setQuizzes(quizzes);
} catch (error) {
    console.error('Failed to load quizzes:', error);
    showError('Could not load quizzes');
}

// DON'T: Let errors propagate silently
const quizzes = await api.getQuizzes();  // BAD: no error handling
```

---

## WebSocket Messages

### Format

All messages are JSON with this structure:

```json
{
    "type": "event:action",
    "payload": { },
    "timestamp": 1234567890
}
```

### Naming

- Type format: `<entity>:<action>`
- Examples: `player:join`, `host:start`, `game:question`

### Go Message Handling

```go
// DO: Explicit type switch
switch msg.Type {
case "player:join":
    handlePlayerJoin(conn, msg.Payload)
case "player:answer":
    handlePlayerAnswer(conn, msg.Payload)
default:
    log.Printf("unknown message type: %s", msg.Type)
}

// DON'T: String parsing
parts := strings.Split(msg, ":")  // BAD
```

---

## Code Style

### Go

- Use `gofmt`. No exceptions.
- Max function length: 50 lines (guideline, not hard rule)
- Max file length: 300 lines (split if larger)
- One struct per file for entities

### TypeScript/Svelte

- Use Prettier with defaults
- Prefer `const` over `let`
- No `any` types unless unavoidable
- Destructure props in Svelte components

---

## Comments

### When to Comment

- **Do:** Explain *why*, not *what*
- **Do:** Document public API functions
- **Don't:** Comment obvious code
- **Don't:** Leave TODO/FIXME (put in TASKS.md instead)

```go
// DO: Explain non-obvious decisions
// roomCode is 6 chars to balance memorability and collision resistance
func generateRoomCode() string {

// DON'T: State the obvious
// GetQuiz gets a quiz
func GetQuiz(id string) (*Quiz, error) {
```

---

## Testing

### Go

- Test files: `*_test.go` in same directory
- Table-driven tests preferred
- Test public functions only (unit tests)

### Frontend

- No tests required for MVP
- If added later: Vitest for unit tests

---

## Git

### Commit Message Format

All commit messages **must** follow this format:

```
<type>(<scope>): <description> [Task X.Y]

[optional body]
```

**Required fields:**
- `type`: One of `feat`, `fix`, `refactor`, `docs`, `chore`, `test`
- `scope`: Affected component (e.g., `quiz`, `ws`, `frontend`)
- `description`: Present tense, max 50 chars, imperative mood
- `[Task X.Y]`: Reference to task in `TASKS.md`

**Examples:**
- `feat(quiz): add delete endpoint [Task 1.3]`
- `fix(ws): handle disconnection race condition [Task 3.2]`
- `docs(domain): add terminology enforcement section [Task 0.1]`

### Commit Frequency

- **At least one commit per task.** Each task must have at least one associated commit.
- **Atomic commits.** One logical change per commit. No bundling unrelated changes.

### Branches

- `main` is always deployable
- Feature branches: `feature/task-X.Y-description`
- No long-lived branches
