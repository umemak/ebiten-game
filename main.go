package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerSize   = 30
)

// Game はゲームの状態を管理する構造体
type Game struct {
	playerX      float64
	playerY      float64
	playerSpeed  float64
	targetX      float64
	targetY      float64
	score        int
	targetActive bool
}

// Update は毎フレーム呼ばれる更新処理
func (g *Game) Update() error {
	// プレイヤーの移動
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerY -= g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerY += g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= g.playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += g.playerSpeed
	}

	// 画面外に出ないように制限
	if g.playerX < 0 {
		g.playerX = 0
	}
	if g.playerX > screenWidth-playerSize {
		g.playerX = screenWidth - playerSize
	}
	if g.playerY < 0 {
		g.playerY = 0
	}
	if g.playerY > screenHeight-playerSize {
		g.playerY = screenHeight - playerSize
	}

	// ターゲットがなければ新しいターゲットを生成
	if !g.targetActive {
		g.targetX = float64(rand.Intn(screenWidth - playerSize))
		g.targetY = float64(rand.Intn(screenHeight - playerSize))
		g.targetActive = true
	}

	// ターゲットとの当たり判定
	if g.playerX < g.targetX+playerSize &&
		g.playerX+playerSize > g.targetX &&
		g.playerY < g.targetY+playerSize &&
		g.playerY+playerSize > g.targetY {
		g.score++
		g.targetActive = false
	}

	// スペースキーでリセット
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.score = 0
	}

	return nil
}

// Draw は毎フレーム呼ばれる描画処理
func (g *Game) Draw(screen *ebiten.Image) {
	// プレイヤーの描画（青い四角）
	ebitenutil.DrawRect(screen, g.playerX, g.playerY, playerSize, playerSize, color.RGBA{0, 0, 255, 255})

	// ターゲットの描画（赤い四角）
	if g.targetActive {
		ebitenutil.DrawRect(screen, g.targetX, g.targetY, playerSize, playerSize, color.RGBA{255, 0, 0, 255})
	}

	// スコアとヘルプの表示
	ebitenutil.DebugPrint(screen, "Score: "+fmt.Sprintf("%d", g.score))
	ebitenutil.DebugPrintAt(screen, "矢印キー: 移動, スペース: リセット", 10, screenHeight-20)
}

// Layout はウィンドウサイズを返す
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Ebiten Game")

	game := &Game{
		playerX:     screenWidth / 2,
		playerY:     screenHeight / 2,
		playerSpeed: 4,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
