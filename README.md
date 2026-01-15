# Quiz Platform

A real-time, multiplayer quiz application that enables hosts to create and manage interactive quiz sessions while players join and compete in real time.

## Overview

Quiz Platform is a web-based application designed for hosting live quiz sessions. A host selects a quiz, starts a game session, and shares a room code with participants. Players join using the code, answer questions within a time limit, and compete for the highest score on the leaderboard.

This project is intended for educational environments, team-building activities, or any scenario where interactive group quizzes are useful. It serves as a portfolio project demonstrating full-stack development with real-time communication.

## Features

- Create and manage quizzes with multiple-choice questions
- Generate unique room codes for game sessions
- Real-time player joining and game state synchronization via WebSocket
- Timed questions with server-side enforcement
- Score tracking and cumulative leaderboard
- Host controls for starting, advancing, and ending games
- Responsive web interface for both hosts and players

## Tech Stack

### Backend

- Go (Golang)
- Fiber (HTTP framework)
- WebSocket (real-time communication)
- MongoDB (persistent quiz storage)

### Frontend

- Svelte 5
- TypeScript
- Vite (build tool)
- Tailwind CSS

### Infrastructure

- Single-server deployment model
- In-memory game state management

## Project Structure

```
quiz-platform/
├── backend/
│   ├── cmd/                # Application entry point
│   └── internal/
│       ├── config/         # Environment configuration
│       ├── entity/         # Domain models
│       ├── handler/        # HTTP and WebSocket handlers
│       ├── repository/     # Data access layer (MongoDB)
│       └── service/        # Business logic
├── frontend/
│   └── src/
│       ├── lib/            # Shared components and utilities
│       └── assets/         # Static assets
├── ARCHITECTURE.md         # System design documentation
├── DOMAIN.md               # Domain terminology reference
├── STYLEGUIDE.md           # Code conventions
└── TASKS.md                # Development task tracking
```

## Installation and Setup

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- MongoDB instance (local or remote)

### Backend Setup

```bash
cd backend
go mod download
go build -o quiz-server ./cmd
```

### Frontend Setup

```bash
cd frontend
npm install
npm run build
```

### Environment Configuration

Create a `.env` file in the backend directory or set the following environment variables:

```
MONGODB_URI=mongodb://localhost:27017/quiz
PORT=3000
CORS_ORIGINS=http://localhost:5173
```

## Usage

### Running the Backend

```bash
cd backend
./quiz-server
```

The server starts on the configured port (default: 3000).

### Running the Frontend (Development)

```bash
cd frontend
npm run dev
```

The development server starts at `http://localhost:5173`.

### Running the Frontend (Production)

Build the frontend and serve the `dist/` directory through the Go backend or a static file server.

### Basic Workflow

1. Host navigates to the application and selects a quiz
2. Host starts a game session and receives a room code
3. Players enter the room code and a nickname to join
4. Host starts the game; questions are displayed to all participants
5. Players submit answers within the time limit
6. Scores are calculated and displayed after each question
7. Final leaderboard is shown at the end of the game

## Inspiration and Ethical Note

This project is inspired by the general concept of interactive, real-time quiz platforms. The implementation is entirely original; no code, design assets, or proprietary logic from any commercial platform was used. This project was developed independently for educational purposes and as a portfolio demonstration of full-stack web development skills.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Disclaimer

This project is not affiliated with, endorsed by, or connected to any commercial quiz platform. All code and design decisions are original work created for educational and portfolio purposes.
