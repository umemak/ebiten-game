# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Running the Game
```bash
go run main.go
```

### Building
```bash
go build
```

### Go Module Management
```bash
go mod tidy    # Clean up dependencies
go mod verify  # Verify dependencies
```

## Project Architecture

This is a simple 2D game built with the Ebiten game library for Go. The architecture follows Ebiten's game loop pattern:

### Core Structure
- **Single file application**: All game logic is contained in `main.go`
- **Game struct**: Manages all game state including player position, target position, score, and game flags
- **Ebiten game loop**: Implements the required `Update()`, `Draw()`, and `Layout()` methods

### Key Components
- **Game State Management**: The `Game` struct holds all mutable state (player coordinates, target coordinates, score, target active flag)
- **Input Handling**: Uses Ebiten's input system for arrow key movement and space key reset
- **Collision Detection**: Simple AABB (Axis-Aligned Bounding Box) collision detection between player and targets
- **Random Target Generation**: Targets spawn at random locations within screen bounds

### Game Loop Flow
1. **Update()**: Processes input, updates game state, handles collisions, manages target spawning
2. **Draw()**: Renders player (blue square), target (red square), score, and controls text
3. **Layout()**: Returns fixed screen dimensions (640x480)

## Code Patterns

### Constants
- Screen dimensions and player size are defined as constants at package level
- All game configuration is centralized in constants section

### Coordinate System
- Uses float64 for precise positioning
- Origin (0,0) is at top-left corner
- Player movement is bounded by screen edges with collision detection

### State Management
- All game state is contained within the Game struct
- No global variables used for game state
- Target spawning is controlled by `targetActive` boolean flag

## Dependencies

The project uses Ebiten v2 (`github.com/hajimehoshi/ebiten/v2`) as the primary game engine. Key modules:
- `ebitenutil`: Utility functions for drawing and debug output
- `inpututil`: Input handling utilities for key press detection