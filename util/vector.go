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

// pos += direction if possible, else returns pos
func SafeMove(pos, direction Vector2D, height, width int) Vector2D {
	if pos.X + direction.X < 0 || pos.X + direction.X >= width {
		return pos
	}
	if pos.Y + direction.Y < 0 || pos.Y + direction.Y >= height {
		return pos
	}
	pos.X += direction.X
	pos.Y += direction.Y
	return pos
}

// pos += direction if possible, else returns false
func (pos *Vector2D) SafeMove(direction Vector2D, height, width int) bool {
	if pos.X + direction.X < 0 || pos.X + direction.X >= width {
		return false
	}
	if pos.Y + direction.Y < 0 || pos.Y + direction.Y >= height {
		return false
	}
	pos.X += direction.X
	pos.Y += direction.Y
	return true
}