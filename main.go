package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/gredinger/pathfinder/controller"
	"github.com/gredinger/pathfinder/model"
	"github.com/gredinger/pathfinder/view"
)

func main() {
	world := model.GenerateBase(100, 100, 0) // 0 is base image, which is water
	world = controller.AgeWorldCycles(world, 15000)
	fmt.Println("The world is aged 15000 years, let's take a look... and check back every 250 years") //
	for i := 0; i < 10000; i += 250 {
		view.GenerateWorld(world, 8, 10000+i)
		world = controller.AgeWorldCycles(world, i)
		fmt.Println("Aging the world a bit more... 250 years.")
	}

	//makeImage()
}

// Generates the entire world
func generateWorld() {
	// Worlds need people
	npc := model.GenerateRandomNPC()

	fmt.Println(npc.Name)
}

func makeImage() {
	gridSize := 128
	gridSpace := 25

	area := gridSpace * gridSize
	dc := gg.NewContext(area, area)
	grass, _ := gg.LoadPNG("assets/grass.png")
	water, _ := gg.LoadPNG("assets/water.png")
	farea := float64(area)
	dc.SetLineWidth(3)
	dc.SetRGB(0, 0, 0)
	for i := 0; i <= area; i += gridSize {
		for y := 0; y <= area; y += gridSize {
			im := water
			if rand.Float64() > .2 {
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

func init() {
	rand.Seed(time.Now().UnixNano())
}
