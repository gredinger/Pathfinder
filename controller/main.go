package controller

import (
	"math/rand"
)

func AgeWorldCycles(world [][]string, age int) [][]string {
	for i := 0; i < age; i++ {
		world = AgeWorld(world)
	}
	return world
}

func AgeWorld(world [][]string) [][]string {
	for x, _ := range world {
		for y, oy := range world[x] {
			if oy == "water" && rand.Uint32() < 100 {
				world = placeTile(x, y, "grass", world)
			}
			if oy == "grass" && rand.Float64() < .007 {
				rx := rand.Intn(2)
				ry := rand.Intn(2)
				if rand.Float64() > .5 {
					world = placeTile(x+rx, y+ry, "grass", world)
				} else {
					world = placeTile(x-rx, y-ry, "grass", world)
				}
			}
			if oy != "water" {
				switch num := rand.Intn(1000000); num {
				case 1:
					world = flood(x, y, rand.Intn(40), world)
				case 100, 200, 300, 400:
					world = flood(x, y, rand.Intn(3), world)
				}
			}
		}
	}
	return world
}

func placeTile(x, y int, tile string, world [][]string) [][]string {
	//fmt.Printf("Placing %v at %v,%v\n", tile, x, y)
	if x < 0 || x >= len(world) {
		return world
	}
	if y < 0 || y >= len(world[x]) {
		return world
	}

	world[x][y] = tile
	return world
}

func flood(x, y, size int, world [][]string) [][]string {
	// fmt.Printf("Fucking your shit around %v,%v with this much water %v\n", x, y, size)
	sx := x - size
	sy := y - size
	for i := 0; i < size; i++ {
		for e := 0; e < size; e++ {
			world = placeTile(sx+i, sy+e, "water", world)
		}
	}
	return world
}
