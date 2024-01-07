package util

type Vector2D struct {
	X int
	Y int
}

func Add(a, b Vector2D) Vector2D {
	return Vector2D{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func MoveUp(start Vector2D) Vector2D {
	return Add(start, Up2D())
}

func MoveDown(start Vector2D) Vector2D {
	return Add(start, Down2D())
}

func MoveLeft(start Vector2D) Vector2D {
	return Add(start, Left2D())
}

func MoveRight(start Vector2D) Vector2D {
	return Add(start, Right2D())
}

func Up2D() Vector2D {
	return Vector2D{0, 1}
}

func Down2D() Vector2D {
	return Vector2D{0, -1}
}

func Left2D() Vector2D {
	return Vector2D{-1, 0}
}

func Right2D() Vector2D {
	return Vector2D{1, 0}
}
