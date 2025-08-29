package geom

type Dims[N Number] Pos[N]

func MakeDimsFromTuple[N Number](x, y N) Dims[N] {
	return Dims[N]{
		X: x,
		Y: y,
	}
}

func (d Dims[N]) ToPos() Pos[N] {
	return Pos[N]{
		X: d.X,
		Y: d.Y,
	}
}

func (d Dims[N]) ToTuple() (N, N) {
	return d.X, d.Y
}

func (d Dims[N]) Eq(other Dims[N]) bool {
	return d.X == other.X && d.Y == other.Y
}

func (d Dims[N]) AlmostEq(other Dims[N], epsilon N) bool {
	return d.ToPos().AlmostEq(other.ToPos(), epsilon)
}

func (d Dims[N]) Scale(factor N) Dims[N] {
	return Dims[N]{
		X: N(float64(d.X) * float64(factor)),
		Y: N(float64(d.Y) * float64(factor)),
	}
}

func (d Dims[N]) Area() N {
	return d.X * d.Y
}
