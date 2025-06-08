# Simple Ebiten Game

A simple 2D game built with the [Ebiten](https://ebiten.org/) game library for Go.

## Game Description

This is a simple target collection game where you control a blue square and try to collect as many red square targets as possible. Each time you collect a target, your score increases and a new target appears at a random location.

## Controls

- **Arrow Keys**: Move the player (blue square)
- **Space**: Reset the score

## Features

- Simple 2D graphics
- Keyboard controls
- Score tracking
- Collision detection
- Random target generation

## Requirements

- Go 1.16 or later
- Ebiten v2

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/umemak/ebiten-game.git
   ```

2. Navigate to the project directory:
   ```
   cd ebiten-game
   ```

3. Run the game:
   ```
   go run main.go
   ```

## Building

To build an executable:

```
go build
```

## Game Mechanics

- The player controls a blue square that can move in four directions
- Red square targets appear at random locations
- When the player collides with a target, the score increases and a new target appears
- The player is constrained to the game window

## Future Improvements

- Add sound effects
- Implement different levels of difficulty
- Add a timer or countdown
- Include obstacles or enemies
- Implement power-ups

## License

This project is open source and available under the [MIT License](LICENSE).
