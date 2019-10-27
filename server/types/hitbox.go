package types

import "fmt"

type Hitbox struct {
	ID       string
	Width    int
	Height   int
	Zindex   int
	Object   *Equipment
	Shape    [][]rune
	Position coordinate // top left corner of the hitbox
}

type Grid struct {
	Width    int
	Height   int
	Entities map[string]*Hitbox
}

type coordinate struct {
	y int
	x int
}

type pixel struct {
	c      rune
	zindex int
}

// input is a Hitbox denoting the arc/shape of an attack. Object is nil in this case
func (g *Grid) DetectCollisions(a *Hitbox) map[string]*Equipment {
	collided := make(map[string]*Equipment)

	for _, e := range g.Entities {
		// coordinates of the attacker relative to the current hitbox being analyzed
		relPos := coordinate{y: a.Position.y - e.Position.y, x: a.Position.x - e.Position.x}

		for y, row := range a.Shape {
			for x, v := range row {
				// coordinates of the current square of the attacker's hibox
				pxPos := coordinate{y: relPos.y + y, x: relPos.x + x}

				// if out of bounds of hitbox
				if pxPos.y < 0 || pxPos.x < 0 || pxPos.y >= e.Height || pxPos.x >= e.Width {
					continue
				}

				if v != ' ' && e.Shape[pxPos.y][pxPos.x] != ' ' {
					collided[e.ID] = e.Object
				}
			}
		}
	}

	return collided
}

func (g *Grid) Draw() {
	grid := make([][]pixel, g.Height, g.Height)

	for i, _ := range grid {
		grid[i] = make([]pixel, g.Width, g.Width)
		for j, _ := range grid[i] {
			grid[i][j] = pixel{c: ' ', zindex: -1}
		}
	}

	for _, e := range g.Entities {
		for y, row := range e.Shape {
			for x, v := range row {
				effy := y + e.Position.y
				effx := x + e.Position.x

				if v != ' ' && e.Zindex > grid[effy][effx].zindex {
					grid[effy][effx].c = v
				}
			}
		}
	}

	for _, row := range grid {
		for _, p := range row {
			fmt.Print(string(p.c))
		}
		fmt.Print("\n")
	}
}
