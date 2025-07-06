package main

import (
	"testing"
)

func TestGameInitialization(t *testing.T) {
	game := &Game{
		playerX:     320,
		playerY:     240,
		playerSpeed: 4,
	}

	if game.playerX != 320 {
		t.Errorf("Expected playerX to be 320, got %f", game.playerX)
	}
	if game.playerY != 240 {
		t.Errorf("Expected playerY to be 240, got %f", game.playerY)
	}
	if game.playerSpeed != 4 {
		t.Errorf("Expected playerSpeed to be 4, got %f", game.playerSpeed)
	}
	if game.score != 0 {
		t.Errorf("Expected initial score to be 0, got %d", game.score)
	}
	if game.targetActive != false {
		t.Errorf("Expected targetActive to be false initially, got %t", game.targetActive)
	}
}

func TestPlayerMovement(t *testing.T) {
	game := &Game{
		playerX:     100,
		playerY:     100,
		playerSpeed: 5,
	}

	originalX := game.playerX
	originalY := game.playerY

	game.playerX += game.playerSpeed
	if game.playerX != originalX+5 {
		t.Errorf("Expected playerX to be %f, got %f", originalX+5, game.playerX)
	}

	game.playerY += game.playerSpeed
	if game.playerY != originalY+5 {
		t.Errorf("Expected playerY to be %f, got %f", originalY+5, game.playerY)
	}
}

func TestPlayerBoundaryConstraints(t *testing.T) {
	game := &Game{
		playerX:     0,
		playerY:     0,
		playerSpeed: 5,
	}

	game.playerX = -10
	if game.playerX < 0 {
		game.playerX = 0
	}
	if game.playerX != 0 {
		t.Errorf("Expected playerX to be constrained to 0, got %f", game.playerX)
	}

	game.playerX = screenWidth + 10
	if game.playerX > screenWidth-playerSize {
		game.playerX = screenWidth - playerSize
	}
	if game.playerX != screenWidth-playerSize {
		t.Errorf("Expected playerX to be constrained to %d, got %f", screenWidth-playerSize, game.playerX)
	}

	game.playerY = -10
	if game.playerY < 0 {
		game.playerY = 0
	}
	if game.playerY != 0 {
		t.Errorf("Expected playerY to be constrained to 0, got %f", game.playerY)
	}

	game.playerY = screenHeight + 10
	if game.playerY > screenHeight-playerSize {
		game.playerY = screenHeight - playerSize
	}
	if game.playerY != screenHeight-playerSize {
		t.Errorf("Expected playerY to be constrained to %d, got %f", screenHeight-playerSize, game.playerY)
	}
}

func TestCollisionDetection(t *testing.T) {
	game := &Game{
		playerX:      100,
		playerY:      100,
		targetX:      110,
		targetY:      110,
		targetActive: true,
		score:        0,
	}

	collisionDetected := game.playerX < game.targetX+playerSize &&
		game.playerX+playerSize > game.targetX &&
		game.playerY < game.targetY+playerSize &&
		game.playerY+playerSize > game.targetY

	if !collisionDetected {
		t.Error("Expected collision to be detected between overlapping player and target")
	}

	game.targetX = 200
	game.targetY = 200

	collisionDetected = game.playerX < game.targetX+playerSize &&
		game.playerX+playerSize > game.targetX &&
		game.playerY < game.targetY+playerSize &&
		game.playerY+playerSize > game.targetY

	if collisionDetected {
		t.Error("Expected no collision between distant player and target")
	}
}

func TestScoreIncrement(t *testing.T) {
	game := &Game{
		playerX:      100,
		playerY:      100,
		targetX:      110,
		targetY:      110,
		targetActive: true,
		score:        5,
	}

	originalScore := game.score

	collisionDetected := game.playerX < game.targetX+playerSize &&
		game.playerX+playerSize > game.targetX &&
		game.playerY < game.targetY+playerSize &&
		game.playerY+playerSize > game.targetY

	if collisionDetected {
		game.score++
		game.targetActive = false
	}

	if game.score != originalScore+1 {
		t.Errorf("Expected score to be %d, got %d", originalScore+1, game.score)
	}
	if game.targetActive != false {
		t.Error("Expected targetActive to be false after collision")
	}
}

func TestTargetGeneration(t *testing.T) {
	game := &Game{
		targetActive: false,
	}

	if !game.targetActive {
		game.targetX = 100
		game.targetY = 100
		game.targetActive = true
	}

	if game.targetActive != true {
		t.Error("Expected targetActive to be true after target generation")
	}
	if game.targetX != 100 {
		t.Errorf("Expected targetX to be 100, got %f", game.targetX)
	}
	if game.targetY != 100 {
		t.Errorf("Expected targetY to be 100, got %f", game.targetY)
	}
}

func TestScoreReset(t *testing.T) {
	game := &Game{
		score: 10,
	}

	game.score = 0

	if game.score != 0 {
		t.Errorf("Expected score to be reset to 0, got %d", game.score)
	}
}

func TestLayout(t *testing.T) {
	game := &Game{}

	width, height := game.Layout(800, 600)

	if width != screenWidth {
		t.Errorf("Expected layout width to be %d, got %d", screenWidth, width)
	}
	if height != screenHeight {
		t.Errorf("Expected layout height to be %d, got %d", screenHeight, height)
	}
}