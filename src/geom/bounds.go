package geom

type Bounds[N Number] struct {
	Pos  Pos[N]
	Dims Dims[N]
}

func MakeBoundsFromPosAndDims[N Number](pos Pos[N], dims Dims[N]) Bounds[N] {
	return Bounds[N]{
		Pos:  pos,
		Dims: dims,
	}
}

func MakeBoundsFromPosAndPos2[N Number](pos Pos[N], pos2 Pos[N]) Bounds[N] {
	return Bounds[N]{
		Pos:  pos,
		Dims: MakeDimsFromTuple(pos2.X-pos.X, pos2.Y-pos.Y),
	}
}

func (b Bounds[N]) OppositeCorner() Pos[N] {
	return Pos[N]{
		X: b.Pos.X + b.Dims.X,
		Y: b.Pos.Y + b.Dims.Y,
	}
}

func (b Bounds[N]) Contains(p Pos[N]) bool {
	p2 := b.OppositeCorner()
	return p.X >= b.Pos.X && p.X < p2.X &&
		p.Y >= b.Pos.Y && p.Y < p2.Y
}
