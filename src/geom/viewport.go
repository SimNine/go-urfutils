package geom

type Viewport[N Number] struct {
	Bounds[N]
	Debug bool
}

func (v *Viewport[N]) ScreenToGame(pos Pos[N]) Pos[N] {
	return pos.TranslatePos(v.Pos)
}

func (v *Viewport[N]) GameToScreen(pos Pos[N]) Pos[N] {
	return pos.Sub(v.Pos)
}
