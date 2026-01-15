# System Architecture

> This document describes the system as it **must** be built. Not aspirationally. Not hypothetically.

---

## System Overview

```
┌──────────────────────────────────────────────────────────┐
│                        CLIENT                            │
│  ┌─────────────────────────────────────────────────────┐ │
│  │              Svelte SPA (Vite + Tailwind)           │ │
│  │  - Host views (dashboard, lobby, game control)      │ │
│  │  - Player views (join, lobby, answer, results)      │ │
│  └─────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────┘
                           │
                           │ HTTP (REST) + WebSocket
                           ▼
┌──────────────────────────────────────────────────────────┐
│                     GO BACKEND                           │
│  ┌──────────────────────┐  ┌───────────────────────────┐ │
│  │     HTTP Handler     │  │   WebSocket Handler       │ │
│  │  - Quiz CRUD         │  │  - Game events            │ │
│  │  - Static files      │  │  - Connection lifecycle   │ │
│  └──────────┬───────────┘  └─────────────┬─────────────┘ │
│             │                            │               │
│  ┌──────────▼────────────────────────────▼─────────────┐ │
│  │                  Service Layer                      │ │
│  │  QuizService     GameService     PlayerService      │ │
│  └──────────┬────────────────────────────┬─────────────┘ │
│             │                            │               │
│  ┌──────────▼───────────┐  ┌─────────────▼─────────────┐ │
│  │   MongoDB (quizzes)  │  │  In-Memory (game state)   │ │
│  └──────────────────────┘  └───────────────────────────┘ │
└──────────────────────────────────────────────────────────┘
```

---

## Included in Scope

| Component | Description |
|-----------|-------------|
| Quiz CRUD | Create, list, delete quizzes via REST |
| Host Flow | Select quiz → Start game → Get room code → Control game |
| Player Flow | Enter code → Set nickname → Wait → Answer → See results |
| Game Flow | Lobby → Question → Collect answers → Show results → Repeat → Leaderboard |
| Leaderboard | Final scores displayed at game end |

---

## Explicitly Out of Scope

| Feature | Reason |
|---------|--------|
| User authentication | Adds complexity, not needed for MVP |
| Persistent game history | Games are ephemeral |
| Quiz editing | Create and delete only; no update |
| Image/media questions | Text only |
| Teams | Individual play only |
| Chat | Not a social platform |
| Mobile apps | Web only |
| Analytics | No tracking |

---

## HTTP vs WebSocket Responsibilities

### HTTP (REST API)

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/api/quizzes` | GET | List all quizzes |
| `/api/quizzes` | POST | Create a quiz |
| `/api/quizzes/:id` | GET | Get single quiz |
| `/api/quizzes/:id` | DELETE | Delete a quiz |

HTTP is for **data that persists**. Quizzes are stored. Games are not.

### WebSocket

| Direction | Event | Purpose |
|-----------|-------|---------|
| Client → Server | `host:create` | Host creates a game room |
| Client → Server | `player:join` | Player joins with code + nickname |
| Client → Server | `host:start` | Host starts the game |
| Client → Server | `host:next` | Host advances to next question |
| Client → Server | `player:answer` | Player submits answer |
| Server → Client | `room:created` | Room code sent to host |
| Server → Client | `player:joined` | Broadcast new player |
| Server → Client | `game:question` | Question data to all |
| Server → Client | `game:results` | Question results to all |
| Server → Client | `game:leaderboard` | Final rankings |

WebSocket is for **real-time, ephemeral events**. No persistence.

---

## Data Storage

### MongoDB (Persistent)

Collection: `quizzes`

```json
{
  "_id": ObjectId,
  "name": "string",
  "questions": [
    {
      "id": "string",
      "text": "string",
      "timeLimit": 20,
      "choices": [
        { "id": "string", "text": "string", "correct": true }
      ]
    }
  ],
  "createdAt": ISODate
}
```

No other collections in MVP.

### In-Memory (Ephemeral)

All game state lives in memory:
- Active game sessions
- Player connections
- Current question index
- Collected answers
- Scores

When the server restarts, all active games are lost. This is acceptable for MVP.

---

## Deployment Model

Single server deployment:

```
┌───────────────────────────────┐
│           Server              │
│  ┌─────────────────────────┐  │
│  │  Go Binary (port 3000)  │  │
│  │  - API + WebSocket      │  │
│  │  - Serves static files  │  │
│  └─────────────────────────┘  │
│                               │
│  ┌─────────────────────────┐  │
│  │  MongoDB (port 27017)   │  │
│  └─────────────────────────┘  │
└───────────────────────────────┘
```

No nginx required for MVP. Go serves everything.

---

## Constraints

1. **Single process.** No worker threads, no background jobs.
2. **Stateful WebSocket.** Game state tied to server memory.
3. **No horizontal scaling.** One server handles all games.
4. **Synchronous operations.** No async job queues.

These constraints are intentional. They keep the system simple and debuggable.

---

## Change Policy

5. **Task-driven changes only.** All architectural changes must be linked to a task in `TASKS.md`.
6. **No implicit refactors.** Do not refactor code, rename packages, or restructure modules unless explicitly defined in a task.
7. **Document before changing.** Architectural modifications require updating `ARCHITECTURE.md` as part of the task scope.
