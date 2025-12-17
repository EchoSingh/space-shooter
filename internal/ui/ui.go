package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

var (
	defaultFont = basicfont.Face7x13
)

// UI handles all UI rendering
type UI struct {
	screenWidth  int
	screenHeight int
}

// NewUI creates a new UI manager
func NewUI(screenWidth, screenHeight int) *UI {
	return &UI{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

// DrawHUD draws the game HUD
func (u *UI) DrawHUD(screen *ebiten.Image, score, health int) {
	// Score
	scoreText := fmt.Sprintf("SCORE: %d", score)
	text.Draw(screen, scoreText, defaultFont, 10, 20, color.White)

	// Health
	healthText := fmt.Sprintf("HEALTH: %d", health)
	text.Draw(screen, healthText, defaultFont, 10, 40, color.White)

	// FPS
	fpsText := fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS())
	text.Draw(screen, fpsText, defaultFont, u.screenWidth-100, 20, color.White)
}

// DrawMenu draws the main menu
func (u *UI) DrawMenu(screen *ebiten.Image) {
	centerX := u.screenWidth / 2
	centerY := u.screenHeight / 2

	// Title
	title := "SPACE SHOOTER"
	titleBounds := text.BoundString(defaultFont, title)
	titleX := centerX - titleBounds.Dx()/2
	text.Draw(screen, title, defaultFont, titleX, centerY-60, color.White)

	// Instructions
	instructions := []string{
		"WASD or Arrow Keys to Move",
		"SPACE to Fire",
		"P to Pause",
		"",
		"Press ENTER to Start",
		"Press ESC to Quit",
	}

	for i, line := range instructions {
		bounds := text.BoundString(defaultFont, line)
		x := centerX - bounds.Dx()/2
		y := centerY + i*20
		text.Draw(screen, line, defaultFont, x, y, color.RGBA{R: 200, G: 200, B: 200, A: 255})
	}
}

// DrawPauseMenu draws the pause menu
func (u *UI) DrawPauseMenu(screen *ebiten.Image) {
	centerX := u.screenWidth / 2
	centerY := u.screenHeight / 2

	// Semi-transparent overlay
	overlay := ebiten.NewImage(u.screenWidth, u.screenHeight)
	overlay.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 128})
	screen.DrawImage(overlay, nil)

	// Pause text
	pauseText := "PAUSED"
	bounds := text.BoundString(defaultFont, pauseText)
	x := centerX - bounds.Dx()/2
	text.Draw(screen, pauseText, defaultFont, x, centerY-20, color.White)

	// Resume instruction
	resumeText := "Press P to Resume"
	resumeBounds := text.BoundString(defaultFont, resumeText)
	resumeX := centerX - resumeBounds.Dx()/2
	text.Draw(screen, resumeText, defaultFont, resumeX, centerY+20, color.RGBA{R: 200, G: 200, B: 200, A: 255})
}

// DrawGameOver draws the game over screen
func (u *UI) DrawGameOver(screen *ebiten.Image, score int) {
	centerX := u.screenWidth / 2
	centerY := u.screenHeight / 2

	// Semi-transparent overlay
	overlay := ebiten.NewImage(u.screenWidth, u.screenHeight)
	overlay.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 128})
	screen.DrawImage(overlay, nil)

	// Game Over text
	gameOverText := "GAME OVER"
	bounds := text.BoundString(defaultFont, gameOverText)
	x := centerX - bounds.Dx()/2
	text.Draw(screen, gameOverText, defaultFont, x, centerY-40, color.RGBA{R: 255, G: 100, B: 100, A: 255})

	// Final score
	scoreText := fmt.Sprintf("Final Score: %d", score)
	scoreBounds := text.BoundString(defaultFont, scoreText)
	scoreX := centerX - scoreBounds.Dx()/2
	text.Draw(screen, scoreText, defaultFont, scoreX, centerY, color.White)

	// Restart instruction
	restartText := "Press ENTER to Restart"
	restartBounds := text.BoundString(defaultFont, restartText)
	restartX := centerX - restartBounds.Dx()/2
	text.Draw(screen, restartText, defaultFont, restartX, centerY+40, color.RGBA{R: 200, G: 200, B: 200, A: 255})

	// Menu instruction
	menuText := "Press ESC for Menu"
	menuBounds := text.BoundString(defaultFont, menuText)
	menuX := centerX - menuBounds.Dx()/2
	text.Draw(screen, menuText, defaultFont, menuX, centerY+60, color.RGBA{R: 200, G: 200, B: 200, A: 255})
}

// DrawText draws text at a position
func (u *UI) DrawText(screen *ebiten.Image, text string, x, y int, col color.Color, fnt font.Face) {
	if fnt == nil {
		fnt = defaultFont
	}
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

// DrawCenteredText draws centered text
func (u *UI) DrawCenteredText(screen *ebiten.Image, str string, y int, col color.Color) {
	bounds := text.BoundString(defaultFont, str)
	x := (u.screenWidth - bounds.Dx()) / 2
	text.Draw(screen, str, defaultFont, x, y, col)
}
