package model

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"

	"github.com/fogleman/gg"
)

type Tile struct {
	Name string
	Im   image.Image
}

var localTiles = []string{
	"water",
	"grass",
}

func GenerateBase(x, y, id int) [][]string {
	tiles := make([][]string, x)
	for i := 0; i < x; i++ {
		tiles[i] = make([]string, y)
		for f := 0; f < y; f++ {
			tiles[i][f] = GetTile(id).Name
		}
	}
	return tiles
}

// Returns the tile by ID, should be a sql select statement dude
func GetTile(id int) Tile {
	return getTiles()[id]
}

// Returns the tile from the tile array you pass to it. Idk how good this is, i feel as though the memory would be duplicated, but at the same time
// I think that Go would compenstate for idiots like me.
func GetTileByName(name string, tiles []Tile) Tile {
	for _, v := range tiles {
		if v.Name == name {
			return v
		}
	}
	return tiles[0]
}

//getTiles is a terrible function that's meant to statically load shit cause we don't have an actual database.
// This should be replaced by like a select all or some shit.
func getTiles() []Tile {
	ta := []Tile{}
	for _, v := range localTiles {
		ta = append(ta, Tile{
			Name: v,
		})
	}
	return ta
}

// LoadsTiles with resizing and with image attached (used to render)
func LoadTilesImageSized(size uint) []Tile {
	tileArray := []Tile{}
	for _, t := range getTiles() {
		tile, _ := gg.LoadPNG(fmt.Sprintf("assets/%s.png", t.Name))
		tR := resize.Resize(size, 0, tile, resize.Lanczos3)
		tileArray = append(tileArray, Tile{
			Name: t.Name,
			Im:   tR,
		})
	}
	return tileArray
}
