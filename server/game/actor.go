package game

// actor states
const (
	actorUpdated int = iota
	actorCreated
	actorDeleted
)

// actor types
const (
	player int = iota
	enemyMeleeBasic
	enemyMeleeFast
	enemyRangedBasic
	enemyRangedAdvanced
	enemyTank
	projectilePlayerBullet
	projectilePlayerHarpoon
	projectileEnemyBullet
	projectileEnemyHarpoon
)

type Actor struct {
	Id            string `json:"id"`
	State         int    `json:"state"`
	Type          int    `json:"type"`
	CurrentHealth int    `json:"currentHealth"`
	maxHealth     int
	Body2D
}

func (a *Actor) takeDamage(amount int) {
	a.CurrentHealth -= amount

	if a.CurrentHealth <= 0 {
		a.State = actorDeleted
	}
}
