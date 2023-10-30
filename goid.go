package main

type Goid struct {
	position     Vector
	velocity     Vector
	acceleration Vector
	maxSpeed     float64
	maxForce     float64
}

func (g *Goid) Seek(t Vector) {
	tv := Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := Sub(tv, g.velocity)
	g.acceleration.Add(force)
}

func (g *Goid) Flee(t Vector) {
	tv := Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := Sub(tv, g.velocity)
	g.acceleration.Sub(force)
}

func (g Goid) IsInsight(g2 Goid) bool {
	d := Sub(g.position, g2.position).Len()
	return d < 100
}
