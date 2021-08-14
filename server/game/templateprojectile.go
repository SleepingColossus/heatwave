package game

type ProjectilesTemplate struct {
	Projectiles []ProjectileTemplate `xml:"projectile"`
}

type ProjectileTemplate struct {
	Type        int     `xml:"type"`
	Damage      int     `xml:"damage"`
	Velocity    float64 `xml:"velocity"`
	MaxDistance float64 `xml:"maxdistance"`
}
