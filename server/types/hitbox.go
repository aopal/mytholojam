package types

type Entity struct {
	ID       string
	Width    int
	Height   int
	Zindex   int
	Object   *Equipment
	Hitbox   [][]rune
	Position coordinate // top left corner of the hitbox
}

type Grid struct {
	Width    int
	Height   int
	Entities map[string]*Entity
}

type coordinate struct {
	y int
	x int
}

// input is an Entity denoting the arc/shape of an attack. Object is nil in this case
func (g *Grid) detectCollisions(a *Entity) map[string]*Equipment {
	collided := make(map[string]*Equipment)

	for _, e := range g.Entities {
		// coordinates of the attacker relative to the current entity being analyzed
		relPos := coordinate{y: a.Position.y - e.Position.y, x: a.Position.x - e.Position.x}

		for y, row := range a.Hitbox {
			for x, v := range row {
				// coordinates of the current square of the attacker's hibox
				pxPos := coordinate{y: relPos.y + y, x: relPos.x + x}

				// if out of bounds of hitbox
				if pxPos.y < 0 || pxPos.x < 0 || pxPos.y >= e.Height || pxPos.x >= e.Width {
					continue
				}

				if v != ' ' && e.Hitbox[pxPos.y][pxPos.x] != ' ' {
					collided[e.ID] = e.Object
				}
			}
		}
	}

	return collided
}

var spearShape [][]rune = [][]rune{
	[]rune{'s', ' '},
	[]rune{'s', 's'},
	[]rune{'s', 's'},
	[]rune{'s', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
	[]rune{' ', 's'},
}
