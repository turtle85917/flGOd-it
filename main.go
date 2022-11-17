package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	screenWidth  = boardSize * tileSize
	screenHeight = boardSize * tileSize
	boardSize    = 10
	tileSize     = 50
)

var colors = [...]color.RGBA{red, orange, yellow, green, blue, purple, brown}
var (
	red    = color.RGBA{255, 0, 0, 255}
	orange = color.RGBA{255, 125, 0, 255}
	yellow = color.RGBA{255, 225, 0, 255}
	green  = color.RGBA{55, 175, 20, 255}
	blue   = color.RGBA{20, 95, 175, 255}
	purple = color.RGBA{85, 20, 175, 255}
	brown  = color.RGBA{110, 60, 20, 255}

	board = [boardSize][boardSize]Tile{}
)

type Tile struct {
	color int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			fx := float64(x * tileSize)
			fy := float64(y * tileSize)
			ebitenutil.DrawRect(screen, fx, fy, tileSize, tileSize, colors[board[y][x].color])
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func reset() {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			board[y][x] = Tile{color: rand.Intn(7)}
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	reset()
}

func main() {
	ebiten.SetWindowTitle("flGOd-it")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
