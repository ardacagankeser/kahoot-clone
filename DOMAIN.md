# Domain Language

> This document defines the **canonical vocabulary** for this project. Use these terms exactly. Do not invent synonyms.

---

## Core Entities

### Quiz

A reusable set of questions created by a host.

| Field | Type | Description |
|-------|------|-------------|
| id | ObjectID | MongoDB identifier |
| name | string | Display name |
| questions | []Question | Ordered list of questions |
| createdAt | timestamp | Creation time |

**Usage:** "Create a quiz", "Delete the quiz", "List quizzes"

**Not:** "Test", "Survey", "Form", "Questionnaire"

---

### Question

A single question within a quiz.

| Field | Type | Description |
|-------|------|-------------|
| id | string | UUID within quiz |
| text | string | The question text |
| timeLimit | int | Seconds allowed (default: 20) |
| choices | []Choice | 2-4 answer options |

**Usage:** "Show the question", "Next question"

**Not:** "Item", "Prompt", "Query"

---

### Choice

An answer option within a question.

| Field | Type | Description |
|-------|------|-------------|
| id | string | UUID within question |
| text | string | Answer text |
| correct | bool | Whether this is the correct answer |

**Usage:** "Select a choice", "The correct choice"

**Not:** "Answer", "Option", "Alternative"

---

### GameSession

A live instance of a quiz being played.

| Field | Type | Description |
|-------|------|-------------|
| id | string | UUID |
| roomCode | string | 6-character join code |
| quiz | Quiz | The quiz being played |
| status | GameStatus | Current phase |
| currentQuestion | int | Index of current question |
| players | map[id]Player | Connected players |
| hostConn | WebSocket | Host connection |

**Usage:** "Start a game session", "End the session"

**Not:** "Game", "Room", "Match", "Lobby" (lobby is a status, not an entity)

---

### Player

A participant in a game session who answers questions.

| Field | Type | Description |
|-------|------|-------------|
| id | string | Connection-based ID |
| nickname | string | Display name (chosen by player) |
| score | int | Cumulative points |
| conn | WebSocket | Player connection |

**Usage:** "Player joined", "Player answered"

**Not:** "User", "Participant", "Attendee", "Client"

---

### Host

The person controlling the game session. Not a stored entity—just the WebSocket connection that created the session.

**Usage:** "Host starts the game", "Host advances question"

**Not:** "Admin", "Teacher", "Creator", "Owner"

---

### Answer

A player's response to a question during a game.

| Field | Type | Description |
|-------|------|-------------|
| playerId | string | Who answered |
| choiceId | string | Which choice selected |
| timestamp | time | When submitted |
| correct | bool | Computed: was it right? |
| points | int | Computed: score earned |

**Usage:** "Submit an answer", "Collect answers"

**Not:** "Response", "Submission", "Reply"

---

## Game Status Values

| Status | Description |
|--------|-------------|
| `LOBBY` | Waiting for players to join |
| `QUESTION` | Question is displayed, timer running |
| `RESULTS` | Question ended, showing results |
| `FINISHED` | Game over, showing final leaderboard |

**Usage:** "The session is in LOBBY status"

---

## Actions (Verbs)

| Verb | Subject | Meaning |
|------|---------|---------|
| create | Quiz, GameSession | Bring into existence |
| delete | Quiz | Remove permanently |
| join | Player → GameSession | Player enters with code |
| leave | Player | Player disconnects |
| start | Host → GameSession | Begin first question |
| advance | Host | Move to next question |
| answer | Player | Submit a choice |
| end | Host, System | Terminate game session |

---

## Naming Conventions

### Variables
- `quiz` not `quizData`, `quizObj`, `q`
- `player` not `p`, `user`, `client`
- `session` not `game`, `room`, `gs`
- `question` not `q`, `ques`, `currentQ`

### Functions
- `CreateQuiz` not `AddQuiz`, `NewQuiz`, `MakeQuiz`
- `JoinSession` not `EnterRoom`, `ConnectToGame`
- `SubmitAnswer` not `SendAnswer`, `PostAnswer`

### Events
- `player:join` not `user:connect`, `client:enter`
- `game:question` not `question:show`, `quiz:question`

---

## Forbidden Terms

Do not use these terms anywhere in code or docs:

| Forbidden | Use Instead |
|-----------|-------------|
| user | player (in game), host (controlling) |
| room | session |
| game | session (the instance), quiz (the template) |
| option | choice |
| response | answer |
| test | quiz |
| admin | host |
