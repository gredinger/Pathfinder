package model

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

type Tile struct {
	Name string
	Im   image.Image
}

//GenerateWorld converts a tile array to image using the gridSpace amount of pixels for each square
func GenerateWorld(world [][]string, gridSpace, year int) {
	dc := gg.NewContext(len(world)*gridSpace, len(world[0])*gridSpace)
	tiles := loadTiles()
	for x := 0; x < len(world); x++ {
		for y := 0; y < len(world[x]); y++ {
			dc.DrawImage(getTile(world[x][y], tiles).Im, x*gridSpace, y*gridSpace)
		}
	}

	dc.SavePNG(fmt.Sprintf("world/%v.png", year))
}

var tiles = []string{
	"water",
	"grass",
}

func GenerateWater(x, y int) [][]string {
	tiles := make([][]string, x)
	for i := 0; i < x; i++ {
		tiles[i] = make([]string, y)
		for f := 0; f < y; f++ {
			tiles[i][f] = "water"
		}
	}
	return tiles
}

func getTile(name string, tiles []Tile) Tile {
	for _, v := range tiles {
		if v.Name == name {
			return v
		}
	}
	return waterTile()
}

func waterTile() Tile {
	tile, _ := gg.LoadPNG("assets/water.png")
	return Tile{
		Name: "water",
		Im:   tile,
	}
}

func loadTiles() []Tile {
	tileArray := []Tile{}
	for _, t := range tiles {
		tile, _ := gg.LoadPNG(fmt.Sprintf("assets/%s.png", t))
		tileArray = append(tileArray, Tile{
			Name: t,
			Im:   tile,
		})
	}
	return tileArray
}
