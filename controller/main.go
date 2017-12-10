package controller

import (
	"fmt"
	"math/rand"
)

func AgeWorld(world [][]string) [][]string {
	for x, _ := range world {
		for y, oy := range world[x] {
			if oy == "water" && rand.Uint32() < 100 {
				world = placeTile(x, y, "grass", world)
			}
			if oy == "grass" && rand.Float64() < .02 {
				rx := rand.Intn(2)
				ry := rand.Intn(2)
				if rand.Float64() > .5 {
					world = placeTile(x+rx, y+ry, "grass", world)
				} else {
					world = placeTile(x-rx, y-ry, "grass", world)
				}
			}
		}
	}
	return world
}

func placeTile(x, y int, tile string, world [][]string) [][]string {
	fmt.Printf("Placing %v at %v,%v\n", tile, x, y)
	if x < 0 || x >= len(world) {
		return world
	}
	if y < 0 || y >= len(world[x]) {
		return world
	}

	world[x][y] = tile
	return world
}
