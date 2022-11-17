package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	image *ebiten.Image
	MinPt image.Point
	x     int
	y     int
	color int
}

func (t *Tile) In(x, y int) bool {
	return t.image.At(x-t.x, y-t.y).(color.RGBA).A > 0
}

func (g *Game) Update() error {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			t := board[y][x]
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				cx, cy := ebiten.CursorPosition()

				idx := y*boardSize + x
				click := cx-t.x <= tileSize*idx && cy-t.y <= tileSize*idx

				if click {

				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			t := board[y][x]
			fx := float64(x * tileSize)
			fy := float64(y * tileSize)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(fx, fy)

			screen.DrawImage(t.image, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func reset() {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			clr := rand.Intn(7)
			tile := ebiten.NewImage(tileSize, tileSize)
			tile.Fill(colors[clr])
			board[y][x] = Tile{image: tile, x: x, y: y, color: clr}
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
