package gameplay

type action struct {
	user   *spirit
	target *targetable // either a spirit or equipment
	move   *move       // name of attacking move, or the special 'swap' move
}
