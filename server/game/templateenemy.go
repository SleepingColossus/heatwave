package game

type EnemiesTemplate struct {
	Enemies []EnemyTemplate `xml:"enemy"`
}

type EnemyTemplate struct {
	Type     int `xml:"type"`
	MaxHP    int `xml:"health"`
	Velocity int `xml:"velocity"`
	Range    int `xml:"range"`
	FireRate int `xml:"firerate"`
}
