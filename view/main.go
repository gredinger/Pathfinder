package view

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/gredinger/pathfinder/model"
)

//GenerateWorld converts a tile array to image using the gridSpace amount of pixels for each square
func GenerateWorld(world [][]string, gridSpace, year int) {
	dc := gg.NewContext(len(world)*gridSpace, len(world[0])*gridSpace)
	tiles := model.LoadTilesImageSized(uint(gridSpace))
	for x := 0; x < len(world); x++ {
		fmt.Printf("World Generation: %v\n", float64(float64(x)/float64(len(world))))
		for y := 0; y < len(world[x]); y++ {
			dc.DrawImage(model.GetTileByName(world[x][y], tiles).Im, x*gridSpace, y*gridSpace)
		}
	}

	dc.SavePNG(fmt.Sprintf("world/%v.png", year))
}
