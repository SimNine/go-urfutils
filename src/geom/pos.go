package geom

import "math"

type Pos[N Number] struct {
	X N
	Y N
}

func MakePosFromTuple[N Number](x, y N) Pos[N] {
	return Pos[N]{
		X: x,
		Y: y,
	}
}

func (p Pos[N]) ToTuple() (N, N) {
	return p.X, p.Y
}

func (p Pos[N]) Translate(dx, dy N) Pos[N] {
	return Pos[N]{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}

func (p Pos[N]) TranslatePos(offset Pos[N]) Pos[N] {
	return p.Translate(offset.X, offset.Y)
}

func (p Pos[N]) Sub(other Pos[N]) Pos[N] {
	return Pos[N]{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Pos[N]) Scale(factor N) Pos[N] {
	return Pos[N]{
		X: N(float64(p.X) * float64(factor)),
		Y: N(float64(p.Y) * float64(factor)),
	}
}

func (p Pos[N]) Eq(other Pos[N]) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Pos[N]) AlmostEq(other Pos[N], epsilon N) bool {
	return math.Abs(float64(p.X-other.X)) < float64(epsilon) &&
		math.Abs(float64(p.Y-other.Y)) < float64(epsilon)
}

func (p Pos[N]) ToFloatPos() Pos[float64] {
	return Pos[float64]{
		X: float64(p.X),
		Y: float64(p.Y),
	}
}
func (p Pos[N]) ToIntPos() Pos[int] {
	return Pos[int]{
		X: int(p.X),
		Y: int(p.Y),
	}
}

func (p Pos[N]) ToDims() Dims[N] {
	return Dims[N](p)
}
