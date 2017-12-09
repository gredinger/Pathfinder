package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/gredinger/pathfinder/model"
)

func main() {
	generateWorld()
	makeImage()
}

// Generates the entire world
func generateWorld() {
	// Worlds need people
	npc := model.GenerateRandomNPC()

	fmt.Println(npc.Name)
}

func makeImage() {
	gridSize := 128
	gridSpace := 100
	area := gridSpace * gridSize
	dc := gg.NewContext(area, area)
	grass, _ := gg.LoadPNG("assets/grass.png")
	water, _ := gg.LoadPNG("assets/water.png")
	farea := float64(area)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetLineWidth(3)
	dc.SetRGB(0, 0, 0)
	for i := 0; i <= area; i += gridSize {
		for y := 0; y <= area; y += gridSize {
			im := water
			if rand.Float64() > .5 {
				im = grass
			}
			dc.DrawImage(im, i, y)
		}
		dc.DrawLine(0, float64(i), farea, float64(i))
		dc.DrawLine(float64(i), 0, float64(i), farea)
	}
	dc.Stroke()
	dc.SavePNG("map5050.png")
}