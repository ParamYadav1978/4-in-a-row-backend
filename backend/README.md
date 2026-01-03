# Backend – 4 in a Row Game

## Purpose
This backend service is responsible for running a real-time multiplayer
4-in-a-Row (Connect Four) game.

The backend handles:
- Player matchmaking
- Game state management
- Turn-based real-time gameplay
- Competitive bot logic
- Player disconnections & reconnections
- Persistence of completed games
- Emission of analytics events

---

## Tech Stack (Decided on Day 1)
- Language: Go (preferred)
- Real-time communication: WebSockets
- Database: PostgreSQL (completed games & leaderboard)
- In-memory store: Active games & players
- Analytics (Bonus): Kafka (event-driven)

---

## High-Level Responsibilities
- Validate every move on the server
- Ensure correct turn order
- Detect win / draw / forfeit
- Assign bot if no opponent joins within 10 seconds
- Allow reconnection within 30 seconds
- different to the bots , based on difficulty.

---

## Notes
This README currently documents system design decisions.
Implementation will begin after design is finalized.
