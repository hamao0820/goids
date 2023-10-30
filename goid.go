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

func (g *Goid) Align(goids []Goid) {
	var avgVel Vector
	n := 0
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		avgVel.Add(other.velocity)
		n++
	}
	if n > 0 {
		avgVel.ScalarMul(1 / float64(n))
		avgVel.Limit(g.maxSpeed)
		g.acceleration.Add(Sub(avgVel, g.velocity))
	}
}

func (g *Goid) Separate(goids []Goid) {
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		d := Sub(g.position, other.position).Len()
		if d < 50 {
			g.Flee(other.position)
		}
	}
}

func (g *Goid) Cohesive(goids []Goid) {
	var avgPos Vector
	n := 0
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		avgPos.Add(other.position)
		n++
	}
	if n > 0 {
		avgPos.ScalarMul(1 / float64(n))
		g.Seek(avgPos)
	}
}

func (g *Goid) Flock(goids []Goid) {
	g.Align(goids)
	g.Separate(goids)
	g.Cohesive(goids)
}

func (g *Goid) AdjustEdge(width, height float64) {
	if g.position.X < 0 {
		g.position.X = 0
		g.velocity.X *= -1
	} else if g.position.X >= width {
		g.position.X = width - 1
		g.velocity.X *= -1
	}

	if g.position.Y < 0 {
		g.position.Y = 0
		g.velocity.Y *= -1
	} else if g.position.Y >= height {
		g.position.Y = height - 1
		g.velocity.Y *= -1
	}
}

func (g *Goid) Update(width, height float64) {
	g.acceleration.Limit(g.maxForce)
	g.velocity.Add(g.acceleration)
	g.velocity.Limit(g.maxSpeed)
	g.position.Add(g.velocity)
	g.acceleration.ScalarMul(0)

	g.AdjustEdge(width, height)
}
