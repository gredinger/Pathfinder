package controller

import (
	"fmt"
	"math/rand"
)

func AgeWorld(world [][]string) {
	for x, _ := range world {
		for y, oy := range world[x] {
			fmt.Printf("%v,%v is %v\n", x, y, oy)
			if oy == "water" && rand.Float64() > .9999 {
				world[x][y] = "grass"
			}
			if oy == "grass" && rand.Float64() > .50 {
				rx := rand.Intn(1)
				ry := rand.Intn(1)
				if rand.Float64() > .5 {
					world[x+rx][y+ry] = "grass"
				} else {
					world[x-rx][y-ry] = "grass"
				}
			}
		}
	}
}
