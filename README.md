# 4 in a Row – Backend Engineering Assignment

This project is a real-time, backend-driven implementation of the
classic 4-in-a-Row (Connect Four) game.

The focus of this project is backend system design, real-time gameplay,
and scalability-oriented thinking.

---

## Game Overview
- Board size: 7 columns x 6 rows
- Players take turns dropping discs into columns
- First player to connect 4 discs (horizontal, vertical, diagonal) wins
- If the board fills with no winner, the game is a draw

---

## System Design (Day 1)

### Tech Stack
- Backend: Go
- Frontend: React (basic)
- Real-time: WebSockets
- Database: PostgreSQL
- Analytics (Bonus): Kafka

---

## High-Level Flow
1. Player enters username
2. Player enters matchmaking queue
3. If another player joins → Player vs Player
4. If no player joins within 10 seconds → Player vs Bot
5. Game starts and proceeds turn-by-turn
6. Backend validates moves and updates game state
7. Game ends (win / draw / forfeit)
8. Result stored in database
9. Leaderboard updated
10. Analytics events emitted

---

## Player Lifecycle
- Idle
- Waiting (matchmaking)
- Playing
- Disconnected (temporary)
- Reconnected (within 30 seconds)
- Finished

---

## Game State Management
- Active games stored in memory
- Completed games stored in PostgreSQL
- Game analytics emitted asynchronously

---

## Status
Day 1 completed: System understanding and design finalized.
