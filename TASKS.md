# Development Tasks

> Ordered, sequential tasks. Complete each before moving to the next. Each task produces a working state.

---

## Status Legend

- `[ ]` Not started
- `[/]` In progress
- `[x]` Complete

---

## Phase 1: Backend Foundation

### Task 1.1: Restructure Backend Folders

**Goal:** Organize code according to STYLEGUIDE.md

**Scope:**
- Rename `collection/` → `repository/`
- Rename `controller/` → `handler/`
- Create `config/` for environment loading
- Move entry point to `cmd/main.go`

**Output:** Backend compiles. Existing endpoints still work.

**Status:** `[ ]`

---

### Task 1.2: Environment Configuration

**Goal:** Load settings from environment, not hardcoded strings

**Scope:**
- Create `internal/config/config.go`
- Load: `MONGODB_URI`, `PORT`, `CORS_ORIGINS`
- Use in `app.go`

**Output:** Server reads config from `.env` or environment variables.

**Status:** `[ ]`

---

### Task 1.3: Quiz CRUD Endpoints

**Goal:** Complete REST API for quiz management

**Scope:**
- `POST /api/quizzes` - Create quiz
- `GET /api/quizzes` - List all (already exists)
- `GET /api/quizzes/:id` - Get single quiz
- `DELETE /api/quizzes/:id` - Delete quiz

**Output:** Can create, list, view, and delete quizzes via curl/Postman.

**Status:** `[ ]`

---

## Phase 2: Game Session Core

### Task 2.1: Game Entities

**Goal:** Define in-memory game state structures

**Scope:**
- Create `entity/game_session.go`
- Create `entity/player.go`
- Create `entity/answer.go`
- Define `GameStatus` enum

**Output:** Structs compile. No usage yet.

**Status:** `[ ]`

---

### Task 2.2: Game Service

**Goal:** Manage game session lifecycle

**Scope:**
- Create `service/game.go`
- `CreateSession(quizID) → roomCode`
- `JoinSession(roomCode, nickname) → player`
- `GetSession(roomCode) → session`
- `DeleteSession(roomCode)`
- In-memory map storage with mutex

**Output:** Unit tests pass for session creation/joining.

**Status:** `[ ]`

---

### Task 2.3: Room Code Generator

**Goal:** Generate unique 6-character codes

**Scope:**
- Alphanumeric, uppercase only (no ambiguous chars: 0/O, 1/I)
- Check for collisions

**Output:** Function generates codes like "A3B7K9".

**Status:** `[ ]`

---

## Phase 3: WebSocket Protocol

### Task 3.1: Message Types

**Goal:** Define all WebSocket message structures

**Scope:**
- Create `ws/message.go`
- Define structs for each event type
- JSON marshaling/unmarshaling

**Output:** Message types documented and usable.

**Status:** `[ ]`

---

### Task 3.2: WebSocket Hub

**Goal:** Manage connections and broadcasting

**Scope:**
- Create `ws/hub.go`
- Track connections by session
- Broadcast to session participants
- Handle disconnections

**Output:** Hub can register connections and broadcast messages.

**Status:** `[ ]`

---

### Task 3.3: Host Flow

**Goal:** Host can create and control a game

**Scope:**
- `host:create` → creates session, returns code
- `host:start` → begins game
- `host:next` → advances question
- `host:end` → ends game

**Output:** Host connects via WebSocket, creates game, sees room code.

**Status:** `[ ]`

---

### Task 3.4: Player Flow

**Goal:** Player can join and participate

**Scope:**
- `player:join` → joins session with nickname
- `player:answer` → submits choice
- Handle disconnection

**Output:** Player joins with code, appears in session.

**Status:** `[ ]`

---

## Phase 4: Game Logic

### Task 4.1: Question Flow

**Goal:** Send questions and manage timing

**Scope:**
- Broadcast question to all participants
- Server-side timer
- Collect answers within time limit
- Reject late answers

**Output:** Questions appear on clients, timer counts down.

**Status:** `[ ]`

---

### Task 4.2: Scoring

**Goal:** Calculate points for correct answers

**Scope:**
- 100 points for correct answer
- No time-based bonus (keep simple)
- Track cumulative score per player

**Output:** Players see their score after each question.

**Status:** `[ ]`

---

### Task 4.3: Leaderboard

**Goal:** Show rankings at end of game

**Scope:**
- Sort players by score
- Broadcast final leaderboard
- Mark session as FINISHED

**Output:** Final screen shows ranked players with scores.

**Status:** `[ ]`

---

## Phase 5: Frontend Structure

### Task 5.1: Routing Setup

**Goal:** SPA navigation between pages

**Scope:**
- Add svelte-routing or hash-based routing
- Define routes: `/`, `/host`, `/host/lobby`, `/host/game`, `/play`, `/play/lobby`, `/play/game`

**Output:** Can navigate between placeholder pages.

**Status:** `[ ]`

---

### Task 5.2: API Service

**Goal:** HTTP client for quiz operations

**Scope:**
- Create `services/api.ts`
- `getQuizzes()`, `createQuiz()`, `deleteQuiz()`
- Error handling

**Output:** Frontend can fetch quiz list.

**Status:** `[ ]`

---

### Task 5.3: WebSocket Service

**Goal:** WebSocket client wrapper

**Scope:**
- Create `services/socket.ts`
- Connect, send, receive
- Reconnection logic (basic)
- Event emitter pattern

**Output:** Can connect to backend WebSocket.

**Status:** `[ ]`

---

### Task 5.4: Game Store

**Goal:** Reactive state for game data

**Scope:**
- Create `stores/game.ts`
- Track: session status, current question, players, score
- Update from WebSocket events

**Output:** Store updates when messages received.

**Status:** `[ ]`

---

## Phase 6: Frontend Screens

### Task 6.1: Home Page

**Goal:** Landing page with role selection

**Scope:**
- "Host a Quiz" button → navigates to /host
- "Join a Game" button + code input → navigates to /play

**Output:** Clean landing page with two clear paths.

**Status:** `[ ]`

---

### Task 6.2: Host Dashboard

**Goal:** Quiz selection for hosting

**Scope:**
- List quizzes from API
- Select quiz → create session
- Navigate to lobby with room code

**Output:** Host sees quizzes, can select one.

**Status:** `[ ]`

---

### Task 6.3: Host Lobby

**Goal:** Waiting room before game starts

**Scope:**
- Display room code prominently
- Show joined players
- "Start Game" button

**Output:** Host sees players joining in real-time.

**Status:** `[ ]`

---

### Task 6.4: Player Join

**Goal:** Enter game code and nickname

**Scope:**
- Input for room code
- Input for nickname
- "Join" button
- Error handling for invalid code

**Output:** Player can enter code and join.

**Status:** `[ ]`

---

### Task 6.5: Player Lobby

**Goal:** Wait for host to start

**Scope:**
- Show "Waiting for host..."
- Display player count

**Output:** Player waits in lobby until game starts.

**Status:** `[ ]`

---

### Task 6.6: Game Screen (Host)

**Goal:** Control game flow and see stats

**Scope:**
- Display current question
- Show how many answered
- "Next Question" button
- Timer display

**Output:** Host controls the game pace.

**Status:** `[ ]`

---

### Task 6.7: Game Screen (Player)

**Goal:** Answer questions

**Scope:**
- Display question text
- 4 answer buttons (colored, Kahoot-style)
- Timer display
- Disable after answering

**Output:** Player can tap to answer.

**Status:** `[ ]`

---

### Task 6.8: Results Screen

**Goal:** Show correct answer and results

**Scope:**
- Reveal correct answer
- Show player's result (correct/wrong)
- Display current score

**Output:** After each question, see results.

**Status:** `[ ]`

---

### Task 6.9: Final Leaderboard

**Goal:** End game with rankings

**Scope:**
- Display ranked list of players
- Highlight top 3
- "Play Again" / "Exit" buttons

**Output:** Game ends with celebratory leaderboard.

**Status:** `[ ]`

---

## Phase 7: Polish & Deploy

### Task 7.1: Error States

**Goal:** Handle failures gracefully

**Scope:**
- Connection lost → show message, offer reconnect
- Invalid room code → show error
- Quiz not found → show error

**Output:** No silent failures.

**Status:** `[ ]`

---

### Task 7.2: Loading States

**Goal:** Show feedback during async operations

**Scope:**
- Loading spinner for quiz list
- "Joining..." while connecting
- Button disabled states

**Output:** UI never feels frozen.

**Status:** `[ ]`

---

### Task 7.3: Visual Polish

**Goal:** Make it look like Kahoot

**Scope:**
- Colored answer buttons (red, blue, green, yellow)
- Large readable fonts
- Animations for transitions
- Mobile-friendly layout

**Output:** Visually appealing, recognizable style.

**Status:** `[ ]`

---

### Task 7.4: Docker Setup

**Goal:** Containerize for deployment

**Scope:**
- Dockerfile for Go backend
- Build frontend, serve as static files
- docker-compose with MongoDB

**Output:** `docker-compose up` runs full stack.

**Status:** `[ ]`

---

### Task 7.5: Production Deploy

**Goal:** Live on a domain

**Scope:**
- Choose platform (Railway/Fly.io/VPS)
- Set environment variables
- Configure domain/SSL
- Test with real users

**Output:** Accessible at https://yourdomain.com

**Status:** `[ ]`

---

## Summary

| Phase | Tasks | Focus |
|-------|-------|-------|
| 1 | 3 | Backend foundation |
| 2 | 3 | Game session core |
| 3 | 4 | WebSocket protocol |
| 4 | 3 | Game logic |
| 5 | 4 | Frontend structure |
| 6 | 9 | Frontend screens |
| 7 | 5 | Polish & deploy |

**Total: 31 tasks**

Estimated time: 2-3 weeks at moderate pace.
