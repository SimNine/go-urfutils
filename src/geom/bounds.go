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

func (b Bounds[N]) Union(other Bounds[N]) Bounds[N] {
	minX := min(b.Pos.X, other.Pos.X)
	minY := min(b.Pos.Y, other.Pos.Y)
	thisOpp := b.OppositeCorner()
	otherOpp := other.OppositeCorner()
	maxX := max(thisOpp.X, otherOpp.X)
	maxY := max(thisOpp.Y, otherOpp.Y)

	return MakeBoundsFromPosAndPos2(
		MakePosFromTuple(minX, minY),
		MakePosFromTuple(maxX, maxY),
	)
}

func (b Bounds[N]) Intersection(other Bounds[N]) (Bounds[N], bool) {
	minX := max(b.Pos.X, other.Pos.X)
	minY := max(b.Pos.Y, other.Pos.Y)
	thisOpp := b.OppositeCorner()
	otherOpp := other.OppositeCorner()
	maxX := min(thisOpp.X, otherOpp.X)
	maxY := min(thisOpp.Y, otherOpp.Y)

	if minX >= maxX || minY >= maxY {
		return Bounds[N]{}, false
	}

	return MakeBoundsFromPosAndPos2(
		MakePosFromTuple(minX, minY),
		MakePosFromTuple(maxX, maxY),
	), true
}
