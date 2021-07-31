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

// TODO consider removing
type Updatable interface {
	update()
}

type Actor struct {
	Id    string `json:"id"`
	State int `json:"state"`
	Type  int `json:"type"`
	Body2D
}
